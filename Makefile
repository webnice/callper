DIR=$(strip $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST)))))

GOPATH := $(DIR):$(GOPATH)
DATE=$(shell date -u +%Y%m%d.%H%M%S.%Z)

default: lint test

test:
	clear
	ln -s . src; true
	GOPATH=${GOPATH} go test -race -cover -v -coverprofile=coverage.log counter
	go test -bench=. -v counter
	go tool cover -html=coverage.log

.PHONY: test

lint:
	gometalinter \
	--vendor \
	--deadline=15m \
	--cyclo-over=10 \
	--disable=aligncheck \
	--skip=src/vendor \
	--linter="vet:go tool vet -printf {path}/*.go:PATH:LINE:MESSAGE" \
	counter/...
.PHONY: lint

clean:
	rm -f ${DIR}/src; true
	rm -rf ${DIR}/bin/*; true
	rm -rf ${DIR}/pkg/*; true
	rm -rf ${DIR}/*.log; true
.PHONY: clean

DIR=$(strip $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST)))))

GOPATH := $(DIR):$(GOPATH)
DATE=$(shell date -u +%Y%m%d.%H%M%S.%Z)

default: lint test

link:
	ln -s . src 2>/dev/null; true
.PHONY: link

test: link
	clear
	GOPATH=${GOPATH} go test -race -cover -v -coverprofile=coverage.log counter
.PHONY: test

cover: test
	GOPATH=${GOPATH} go tool cover -html=coverage.log
.PHONY: cover

bench: link
	GOPATH=${GOPATH} go test -race -bench=. ./...
.PHONY: bench

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

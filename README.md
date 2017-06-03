# callper

[![GoDoc](https://godoc.org/github.com/webnice/callper?status.svg)](https://godoc.org/github.com/webnice/callper)
[![Coverage Status](https://coveralls.io/repos/github/webnice/callper/badge.svg?branch=master)](https://coveralls.io/github/webnice/callper?branch=master)

Golang library

Calculating the current percentage of Tic() calls relative to previous calls in the last 5 minutes.

#### Dependencies

	NONE

#### Installation
```bash
go get gopkg.in/webnice/callper.v1/counter
```

#### Usage
```golang
import "gopkg.in/webnice/callper.v1/counter"

var percent float64
var c = counter.new()

c.Tic()
percent = c.Percent()
```

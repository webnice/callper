# callper

[![GoDoc](https://godoc.org/gopkg.in/webnice/callper.v1/counter?status.svg)](https://godoc.org/gopkg.in/webnice/callper.v1/counter)
[![Coverage Status](https://coveralls.io/repos/github/webnice/callper/badge.svg?branch=v1)](https://coveralls.io/github/webnice/callper?branch=v1)
[![Build Status](https://travis-ci.org/webnice/callper.svg?branch=v1)](https://travis-ci.org/webnice/callper)
[![CircleCI](https://circleci.com/gh/webnice/callper/tree/v1.svg?style=svg)](https://circleci.com/gh/webnice/callper/tree/v1)

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

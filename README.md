# callper

[![GoDoc](https://godoc.org/github.com/webnice/callper?status.svg)](https://godoc.org/github.com/webnice/callper)

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
percent.Percent()
```

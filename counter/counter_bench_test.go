package counter

import (
	"testing"
	"time"
)

var beco Interface

func init() {
	beco = New()
	go func() {
		var tic *time.Ticker
		var c int64
		tic = time.NewTicker(time.Nanosecond)
		for {
			<-tic.C
			c++
			beco.Hit()
			if c > 1000000000 {
				break
			}

		}
	}()
	time.Sleep(time.Second / 2)
}

func BenchmarkTic(b *testing.B) {
	for n := 0; n < b.N; n++ {
		beco.Hit()
	}
}

func BenchmarkClean(b *testing.B) {
	var gist = beco.(*impl)
	for n := 0; n < b.N; n++ {
		gist.Clean()
	}
}

func BenchmarkPercent(b *testing.B) {
	var gist *impl
	var key int64
	var count uint32

	// Info
	gist = beco.(*impl)
	gist.RLock()
	for key = range gist.mem {
		count += gist.mem[key]
	}
	gist.RUnlock()

	for n := 0; n < b.N; n++ {
		beco.Percent()
	}
}

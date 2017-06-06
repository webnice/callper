package counter

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	cou := New().(*impl)
	if cou.mem == nil {
		t.Errorf("Error in New()")
	}
	if cou.Percent() != 100 {
		t.Errorf("Error in Percent()")
	}
}

func TestClean(t *testing.T) {
	var gist *impl
	var tic, end *time.Ticker
	var count float64
	var key int64
	var ok bool

	gist = New().(*impl)
	gist.averageCount = 3
	gist.averageDuration = time.Second
	tic = time.NewTicker(time.Second / 2)
	end = time.NewTicker(time.Second * 10)

	for {
		if ok {
			break
		}
		select {
		case <-tic.C:
			gist.Tic()
		case <-end.C:
			ok = true
		}
	}
	defer tic.Stop()
	defer end.Stop()

	// Count
	for key = range gist.mem {
		if time.Since(time.Unix(0, key)) <= gist.averageDuration*time.Duration(gist.averageCount) {
			count++
		}
	}
	if count != 6 {
		t.Errorf("Clean() error")
	}
}

func TestPercent(t *testing.T) {
	var gist *impl
	var key int64
	var i int

	gist = New().(*impl)
	gist.averageCount = 5
	gist.averageDuration = time.Minute

	for i = 1; i <= 10; i++ {
		key = time.Now().Add(0 - time.Minute*time.Duration(10-i-1)*2).Truncate(gist.averageDuration / 10).UnixNano()
		gist.mem[key] = float64(i * 2)
	}
	if gist.Percent() != check(gist.mem, gist.averageDuration, gist.averageCount) {
		t.Error("Error Percent()")
	}
}

func check(mem map[int64]float64, d time.Duration, n float64) (percent float64) {
	var tm time.Time
	var du time.Duration
	var key int64
	var sum float64
	var cur float64

	for key = range mem {
		tm = time.Unix(0, key)
		du = time.Since(tm)
		if du <= 0 || du < d {
			cur += mem[key]
		} else {
			sum += mem[key]
		}
	}
	percent = sum / (n - 1)
	percent = cur / percent * 100

	return
}

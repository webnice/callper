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
	var count uint32
	var key int64
	var ok bool

	gist = New().(*impl)
	gist.count = 3
	gist.duration = time.Second
	tic = time.NewTicker(time.Second / 100)
	end = time.NewTicker(time.Second * 10)

	for {
		if ok {
			break
		}
		select {
		case <-tic.C:
			gist.Hit()
		case <-end.C:
			ok = true
		}
	}
	defer tic.Stop()
	defer end.Stop()

	// Count
	for key = range gist.mem {
		if time.Since(time.Unix(0, key)) <= gist.duration*time.Duration(gist.count) {
			count++
		}
	}
	if count != 3 {
		t.Errorf("Clean() error, count is %d expected 3", count)
	}
}

func TestPercent(t *testing.T) {
	var gist *impl
	var key int64
	var percent float64
	var i int

	gist = New().NewCounts(10).NewDuration(time.Second / 2).(*impl)
	for i = 1; i <= 20; i++ {
		key = time.Now().Add(0 - gist.duration*time.Duration(10-i-1)*2).Truncate(gist.duration / 10).UnixNano()
		gist.mem[key] = uint32(i * 2)
	}
	percent = gist.Percent()
	if int64(percent) != 142 {
		t.Errorf("Error Percent(), return %f expected 142.767857", percent)
	}
}

package counter

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
import (
	"sort"
	"testing"
	"time"
)

type sortReverseKeys []int64

func (s sortReverseKeys) Len() int           { return len(s) }
func (s sortReverseKeys) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortReverseKeys) Less(i, j int) bool { return s[i] > s[j] }

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
	if gist.Percent() != check(gist.mem, gist.averageCount) {
		t.Error("Error Percent()")
	}
}

func check(mem map[int64]float64, n float64) (percent float64) {
	var key int64
	var keys []int64
	var sum float64

	for key = range mem {
		keys = append(keys, key)
	}
	sort.Sort(sortReverseKeys(keys))
	for key = range mem {
		if key != keys[0] {
			sum += mem[key]
		}
	}
	percent = sum / (n - 1)
	percent = mem[keys[0]] / percent * 100

	return
}

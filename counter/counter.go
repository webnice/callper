package counter

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
import (
	"time"
)

// New creates a new object and return interface
func New() Interface {
	var cou = new(impl)
	cou.mem = make(map[int64]uint32)
	cou.duration = defaultDuration
	cou.count = defaultCounts
	return cou
}

// NewDuration Set new interval duration of the account of hits
func (cou *impl) NewDuration(duration time.Duration) Interface { cou.duration = duration; return cou }

// NewCounts Set new amount of intervals for recording hits history
func (cou *impl) NewCounts(count uint64) Interface { cou.count = count; return cou }

// Hit Increment of the hit in current minutes
func (cou *impl) Hit() Interface {
	var key int64
	var ok bool

	key = time.Now().Truncate(cou.duration).UnixNano()
	cou.Lock()
	if _, ok = cou.mem[key]; ok {
		cou.mem[key]++
	} else {
		cou.mem[key] = 1
	}
	cou.Unlock()
	cou.Clean()
	return cou
}

// Clean Очистка от старых данных
func (cou *impl) Clean() {
	var key int64
	var del []int64
	var tm time.Time

	cou.RLock()
	for key = range cou.mem {
		tm = time.Unix(0, key)
		if time.Since(tm) > cou.duration*time.Duration(cou.count) {
			del = append(del, key)
		}
	}
	cou.RUnlock()
	cou.Lock()
	for _, key = range del {
		delete(cou.mem, key)
	}
	cou.Unlock()
}

func (cou *impl) history() (ret []uint32) {
	var key int64
	var value uint32
	var tm time.Time
	var du time.Duration
	var n uint64

	ret = make([]uint32, cou.count)
	cou.RLock()
	for key, value = range cou.mem {
		tm = time.Unix(0, key)
		du = time.Since(tm)
		n = uint64(du / cou.duration)
		if n < cou.count {
			ret[n] += value
		}
	}
	cou.RUnlock()

	return
}

// Percent Процент обращений текущего интервала по отношению к предыдущим
func (cou *impl) Percent() (ret float64) {
	var history []uint32
	var n, m uint64
	var ap float64

	cou.Clean()
	history = cou.history()

	for n = 1; n < cou.count; n++ {
		if history[n] == 0 {
			continue
		}
		ap += float64(history[0]) / float64(history[n])
		m++
	}
	if m == 0 {
		ret = 100
		return
	}
	ret = ap / float64(m) * 100

	return
}

// IsFirst Returns true if in first interval zero hits
func (cou *impl) IsFirst() (ret bool) {
	var history []uint32

	cou.Clean()
	history = cou.history()
	if history[0] == 0 {
		ret = true
	}

	return
}

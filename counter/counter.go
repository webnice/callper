package counter

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
import (
	"time"
)

// New creates a new object and return interface
func New() Interface {
	var cou = new(impl)
	cou.mem = make(map[int64]float64)
	cou.averageCount = _AverageCount
	cou.averageDuration = _AverageDuration
	return cou
}

// Tic Счётчик
func (cou *impl) Tic() {
	var key int64
	var ok bool

	key = time.Now().Truncate(cou.averageDuration / 10).UnixNano()
	cou.Lock()
	if _, ok = cou.mem[key]; ok {
		cou.mem[key]++
	} else {
		cou.mem[key] = 1
	}
	cou.Unlock()
	cou.Clean()
}

// Clean Очистка от старых данных
func (cou *impl) Clean() {
	var key int64
	var del []int64
	var tm time.Time

	cou.RLock()
	for key = range cou.mem {
		tm = time.Unix(0, key)
		if time.Since(tm) > cou.averageDuration*time.Duration(cou.averageCount) {
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

// Percent Процент обращений текущей минуты по отношению к предыдущим
func (cou *impl) Percent() (ret float64) {
	var tm time.Time
	var du time.Duration
	var key int64
	var pre, cur float64

	cou.Clean()
	cou.RLock()
	for key = range cou.mem {
		tm = time.Unix(0, key)
		du = time.Since(tm)
		if du <= 0 || du < cou.averageDuration {
			cur += cou.mem[key]
		} else {
			pre += cou.mem[key]
		}
	}
	cou.RUnlock()
	pre = pre / (cou.averageCount - 1)
	if pre == 0 {
		ret = 100
	} else {
		ret = cur / pre * 100
	}

	return
}

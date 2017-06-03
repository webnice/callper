package counter

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
import (
	"sync"
	"time"
)

const (
	_AverageDuration time.Duration = time.Minute
	_AverageCount    float64       = 5
)

// Interface is an interface of package
type Interface interface {
	// Tic Счётчик
	Tic()

	// Percent Процент обращений текущей минуты по отношению к предыдущим
	Percent() float64
}

// impl is an implementation of package
type impl struct {
	sync.RWMutex
	mem             map[int64]float64
	averageDuration time.Duration
	averageCount    float64
}

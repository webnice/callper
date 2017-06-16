package counter

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
import (
	"sync"
	"time"
)

const (
	defaultDuration time.Duration = time.Minute
	defaultCounts   uint64        = 5
)

// Interface is an interface of package
type Interface interface {
	// Hit Increment of the hit in current minutes
	Hit() Interface

	// Percent The percentage of hits of the current minute in relation to the average number of hits for the entire history of all intervals
	Percent() float64

	// IsFirst Returns true if in first interval zero hits
	IsFirst() bool

	// NewDuration Set new interval duration of the account of hits
	NewDuration(duration time.Duration) Interface

	// NewCounts Set new amount of intervals for recording hits history
	NewCounts(count uint64) Interface
}

// impl is an implementation of package
type impl struct {
	sync.RWMutex
	mem      map[int64]uint32
	duration time.Duration
	count    uint64
}

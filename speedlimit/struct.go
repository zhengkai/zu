package speedlimit

import (
	"sync"
	"time"
)

type limiter struct {
	duration time.Duration
	max      int
	mux      sync.Mutex
	history  []uint64
	count    int
}

func (li *limiter) CanAdd() (ok bool) {

	now := time.Now()

	li.mux.Lock()

	if li.count < li.max {
		ok = true
		go li.add(now)
		return
	}

	var i int
	var v uint64
	threshold := uint64(now.Add(-li.duration).UnixNano() / 1000000)
	for i, v = range li.history {
		if v >= threshold {
			break
		}
	}
	if i == 0 {
		li.mux.Unlock()
		return
	}

	ok = true

	go li.renewHistory(now, i)

	return
}

func (li *limiter) add(now time.Time) {

	ms := uint64(now.UnixNano() / 1000000)
	li.count++
	li.history = append(li.history, ms)

	li.mux.Unlock()
}

func (li *limiter) renewHistory(now time.Time, i int) {

	ms := uint64(now.UnixNano() / 1000000)

	count := li.count - i

	history := make([]uint64, count)
	if count > 1 {
		copy(history, li.history[i:])
	}
	history[count-1] = ms
	li.history = history
	li.count = count

	li.mux.Unlock()
}

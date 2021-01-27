package speedlimit

import "time"

// New create a limiter
func New(d time.Duration, max int) (li Limiter, err error) {

	a := &limiter{
		max:      max,
		duration: d,
	}

	return a, nil
}

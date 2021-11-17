package conlimit

import (
	"sync"
)

// Limiter ...
type Limiter struct {
	mux      sync.Mutex
	count    int
	capacity int
}

// New ...
func New(capacity int) *Limiter {
	return &Limiter{
		capacity: capacity,
	}
}

// Add ...
func (c *Limiter) Add() (ok bool) {

	if c.count >= c.capacity {
		return
	}

	c.mux.Lock()
	defer c.mux.Unlock()
	if c.count >= c.capacity {
		return
	}

	c.count++
	ok = true

	return
}

// Done ...
func (c *Limiter) Done() {
	c.mux.Lock()
	c.count--
	c.mux.Unlock()
}

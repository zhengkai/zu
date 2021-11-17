package conlimit

import (
	"fmt"
	"sync"
)

// Con ...
type Con struct {
	mux      sync.Mutex
	count    int
	capacity int
}

// New ...
func New(capacity int) *Con {
	return &Con{
		capacity: capacity,
	}
}

// Add ...
func (c *Con) Add() (ok bool) {

	fmt.Println(c.count)

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
func (c *Con) Done() {
	c.mux.Lock()
	c.count--
	c.mux.Unlock()
}

package common

import (
	"runtime"
	"sync"
)

// Actor is executable function
type Actor = func()

// Pool is groutine pool
type Pool struct {
	count    int             // goroutine count 
	recv     chan Actor      // receive actor
	wg       *sync.WaitGroup // to wait all goroutine complete
}

// NewPool creates a new pool with [ count <= 0 ? CPU Num : count ] miles
func NewPool(count int) *Pool {
	if count <= 0 {
		count = runtime.NumCPU()
	} 
	pool := &Pool{
		count:    count, 
		recv:     make(chan Actor, count),
		wg:       &sync.WaitGroup{},
	}
	pool.wg.Add(count)
	for i := 0; i < count; i++ { 
		go func() {
			for actor := range pool.recv {
				actor()
			}
			pool.wg.Done()
		}() 
	}  
	return pool
}

// Go run a actor in pool
func (p *Pool) Go(actor Actor) {
	p.recv <- actor
}

// Close the pool
func (p *Pool) Close() {
	close(p.recv) 
}

// CloseAndWaitAll close the pool and wait all goroutine completed
func (p *Pool) CloseAndWaitAll() {
	p.Close()
	p.wg.Wait()
}

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
	channels []chan Actor    //
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
		channels: []chan Actor{},
		recv:     make(chan Actor, count),
		wg:       &sync.WaitGroup{},
	}
	pool.wg.Add(count)
	for i := 0; i < count; i++ {
		ch := make(chan Actor)
		go func() {
			for actor := range ch {
				actor()
			}
			pool.wg.Done()
		}()
		pool.channels = append(pool.channels, ch)
	}
	go func() {
		i := 0
		for actor := range pool.recv {
			for {
				select {
				case pool.channels[i] <- actor:
					break
				default:
				}
				i++
				if i == pool.count {
					i = 0
				}
			}
		}
	}()
	return pool
}

// Go run a actor in pool
func (p *Pool) Go(actor Actor) {
	p.recv <- actor
}

// Close the pool
func (p *Pool) Close() {
	close(p.recv)
	for _, c := range p.channels {
		close(c)
	}
}

// CloseAndWaitAll close the pool and wait all goroutine completed
func (p *Pool) CloseAndWaitAll() {
	p.Close()
	p.wg.Wait()
}

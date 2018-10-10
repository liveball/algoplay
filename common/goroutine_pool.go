package common

import (
	"runtime"
	"sync"
)

// Actor is executable function
type Actor = func()

type Pool struct {
	count    int             // goroutine count
	channels []chan Actor    //
	recv     chan Actor      // receive actor
	wg       *sync.WaitGroup // to wait all goroutine complete
}

func NewPool(count int) *Pool {
	if count <= 0 {
		count = runtime.NumCPU()
	}
	recv := make(chan Actor, count)
	pool := &Pool{
		count:    count,
		channels: []chan Actor{},
		recv:     recv,
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
		for actor := range recv {
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

func (p *Pool) Go(actor Actor) {
	p.recv <- actor
}

func (p *Pool) Close() {
	close(p.recv)
	for _, c := range p.channels {
		close(c)
	}
}

func (p *Pool) CloseAndWaitAll() {
	p.Close()
	p.wg.Wait()
}

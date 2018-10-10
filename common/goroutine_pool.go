package common

import (
	"runtime"
	"sync"
)

// Actor is executable function
type Actor = func()

// Pool is groutine pool
type Pool struct {
	count   int            // goroutine count
	recv    chan Actor     // receive actor
	chLimit chan struct{}  //limit chan buffer
	wg      sync.WaitGroup // to wait all goroutine complete
}

//New for pool instance
func New(count int) *Pool {
	if count <= 0 {
		// runtime.GOMAXPROCS(1)
		count = runtime.NumCPU()
	}

	recv := make(chan Actor, count)
	p := Pool{
		count:   count,
		recv:    recv,
		chLimit: make(chan struct{}, 1),
	}

	p.wg.Add(count)
	for i := 0; i < count; i++ {
		go func() {
			for actor := range p.recv {
				p.chLimit <- struct{}{}
				actor()
			}
			p.wg.Done()
		}()
	}

	return &p
}

//Go send func to chan.
func (p *Pool) Go(actor Actor) {
	p.recv <- actor
	<-p.chLimit
}

//Close close recv chan
func (p *Pool) Close() {
	close(p.recv)
	p.wg.Wait()
}

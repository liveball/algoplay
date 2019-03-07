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
func New(numGoroutines int) *Pool {
	if numGoroutines < 1 {
		numGoroutines = runtime.NumCPU()
	}

	runtime.GOMAXPROCS(numGoroutines) //set P num P<=2 有序 P>2 无序

	p := Pool{
		count:   numGoroutines,
		recv:    make(chan Actor, numGoroutines),
		chLimit: make(chan struct{}, 1),
	}

	fn := func() {
		for actor := range p.recv {
			p.chLimit <- struct{}{}
			actor()
		}
		p.wg.Done()
	}

	p.wg.Add(numGoroutines)
	if numGoroutines == 1 {
		go fn()
	} else {
		for i := 0; i < numGoroutines; i++ {
			go fn()
		}
	}

	return &p
}

//Go send func to chan.
func (p *Pool) Go(actor Actor) {
	go func() {
		p.recv <- actor
		<-p.chLimit
	}()
}

//Close close recv chan
func (p *Pool) Close() {
	close(p.recv)
	p.wg.Wait()
}

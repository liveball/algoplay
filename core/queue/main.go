package main

import (
	"fmt"

	"algoplay/core/queue/lock"
	"algoplay/core/queue/lockless"
)

func main() {
	//testLock()
	testLockLess()
}

func testLockLess() {
	q, err := lockless.New(10)
	if err != nil {
		fmt.Printf("lockless.New error(%v)\n", err)
		return
	}
	for i := 0; i < 300; i++ {
		q.Push(i)
		code, v := q.Get()
		if code != 0 {
			fmt.Printf("q.Get code(%v)\n", code)
		}
		println(111, v.(int))
	}
}

func testLock() {
	var q lock.Queue

	for i := 0; i < 3; i++ {
		q.Enqueue(i + 1)
		v, ok := q.Dequeue()
		if !ok {
			break
		}
		println(v)
	}

	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 100; i++ {
				q.Enqueue(i + 1)
			}
		}()

		go func() {
			for {
				v, ok := q.Dequeue()
				if !ok {
					break
				}
				_ = v
				// println(v)
			}
		}()
	}
}

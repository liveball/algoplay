package simple

import (
	"fmt"
	"log"
)

type ArrayQueue struct {
	head int
	tail int
	cap  int
	q    []interface{}
}

func NewArrayQueue(cap int) (aq *ArrayQueue) {
	aq = &ArrayQueue{
		head: 0,
		tail: 0,
		cap:  cap,
		q:    make([]interface{}, cap),
	}

	return
}

func (aq *ArrayQueue) EnQueue(v interface{}) bool {
	if aq.tail == aq.cap {
		log.Println("queue is full")
		return false
	}

	aq.q[aq.tail] = v

	aq.tail++

	return true
}

func (aq *ArrayQueue) DeQueue() interface{} {
	if aq.head == aq.tail {
		log.Println("queue is empty")
		return nil
	}

	v := aq.q[aq.head]

	aq.head++

	return v
}

func (aq *ArrayQueue) Length() int {
	return aq.tail - aq.head
}

func (aq *ArrayQueue) String() string {
	if aq.head == aq.tail {
		return "empty queue"
	}

	result := "head"
	for i := aq.head; i <= aq.tail-1; i++ {
		result += fmt.Sprintf("<-%+v", aq.q[i])
	}

	result += "<-tail"
	return result
}

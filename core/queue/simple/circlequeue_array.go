package simple

import (
	"fmt"
	"log"
)

type ArrayCircleQueue struct {
	head int
	tail int
	cap  int
	q    []interface{}
}

func NewArrayCircleQueue(cap int) (aq *ArrayCircleQueue) {
	aq = &ArrayCircleQueue{
		head: 0,
		tail: 0,
		cap:  cap,
		q:    make([]interface{}, cap),
	}

	return
}

func (aq *ArrayCircleQueue) EnQueue(v interface{}) bool {
	if aq.IsFull() {
		log.Println("queue is full")
		return false
	}

	aq.q[aq.tail] = v

	aq.tail = (aq.tail + 1) % aq.cap

	return true
}

func (aq *ArrayCircleQueue) DeQueue() interface{} {
	if aq.IsEmpty() {
		log.Println("queue is empty")
		return nil
	}

	v := aq.q[aq.head]

	aq.head = (aq.head + 1) % aq.cap

	return v
}

func (aq *ArrayCircleQueue) IsFull() bool {
	if aq.head == (aq.tail+1)%aq.cap {
		return true
	}

	return false
}

func (aq *ArrayCircleQueue) IsEmpty() bool {
	if aq.tail == aq.head {
		return true
	}

	return false
}

func (aq *ArrayCircleQueue) Length() int {
	return aq.tail - aq.head
}

func (aq *ArrayCircleQueue) String() string {
	if aq.IsEmpty() {
		return "empty queue"
	}

	result := "head"
	var i = aq.head

	for {
		result += fmt.Sprintf("<-%+v", aq.q[i])
		i = (i + 1) % aq.cap

		if i == aq.tail {
			break
		}
	}

	result += "<-tail"
	return result
}

package simple

import (
	"fmt"
	"log"
)

type Node struct {
	next *Node
	val  interface{}
}

type LinkQueue struct {
	head   *Node
	tail   *Node
	length int
	cap    int
}

func NewLinkQueue(cap int) (lq *LinkQueue) {
	lq = &LinkQueue{
		head:   nil,
		tail:   nil,
		length: 0,
		cap:    cap,
	}

	return
}

func (lq *LinkQueue) EnQueue(val interface{}) bool {
	if lq.length == lq.cap {
		log.Println("queue is full")
		return false
	}

	node := &Node{
		next: nil,
		val:  val,
	}

	if lq.tail == nil { //如果队尾指针为空，则队列为空，队头和队尾指针同时指向新节点
		lq.head = node
		lq.tail = node
	} else {
		lq.tail.next = node //队尾新增节点
		lq.tail = node      //新节点作为新队尾
	}

	lq.length++ //每入队一个元素队列长度加1

	return true
}

func (lq *LinkQueue) DeQueue() interface{} {
	if lq.head == nil {
		log.Println("queue is empty")
		return nil
	}

	v := lq.head.val //获取队头节点数据

	lq.head = lq.head.next //队头后移一个节点

	lq.length-- //每出队一个元素队列长度减1

	return v
}

func (lq *LinkQueue) Length() int {
	return lq.length
}

func (lq *LinkQueue) String() string {
	if lq.head == nil {
		return "empty queue"
	}

	result := "head"
	for cur := lq.head; cur != nil; cur = cur.next {
		result += fmt.Sprintf("<-%+v", cur.val)
	}
	result += "<-tail"

	return result
}

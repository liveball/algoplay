package simple

import (
	"testing"
)

func Test_linkList_EnQueue(t *testing.T) {
	q := NewLinkQueue(3)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	q.EnQueue(5)
	q.EnQueue(6)
	t.Log(q)
}

func Test_linkList_DeQueue(t *testing.T) {
	q := NewLinkQueue(6)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	q.EnQueue(5)
	q.EnQueue(6)

	t.Log(q) //1-6
	q.DeQueue()
	t.Log(q)
	q.DeQueue()
	t.Log(q)
	q.DeQueue()
	t.Log(q)
	q.DeQueue()
	t.Log(q)
	q.DeQueue()
	t.Log(q)

	q.DeQueue()
	t.Log(q) //empty queue
}

func Benchmark_linkList_EnQueue(b *testing.B) {
	lq := NewLinkQueue(50000)

	for i := 0; i < b.N; i++ {
		lq.EnQueue(i)
		//fmt.Println("入队：", i, lq.EnQueue(i))
	}

	//for lq.Length() != 0 {
	//	fmt.Println("出队：", lq.DeQueue())
	//}
}

func Test_array_enQueue(t *testing.T) {
	q := NewArrayQueue(5)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	q.EnQueue(5)
	q.EnQueue(6)
	t.Log(q)
}

func Test_array_DeQueue(t *testing.T) {
	q := NewArrayQueue(5)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	t.Log(q)
	q.DeQueue()
	t.Log(q)
	q.DeQueue()
	t.Log(q)
	q.DeQueue()
}

func TestCircularQueue_EnQueue(t *testing.T) {
	q := NewArrayCircleQueue(5)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	q.EnQueue(5)
	q.EnQueue(6)
	t.Log(q)
}

func TestCircularQueue_DeQueue(t *testing.T) {
	q := NewArrayCircleQueue(5)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	q.EnQueue(5)
	q.EnQueue(6)
	t.Log(q)
	t.Log(q.DeQueue())
	t.Log(q)
	q.EnQueue(5)
	t.Log(q)
	q.DeQueue()
	t.Log(q)
	q.DeQueue()
	t.Log(q)
	q.DeQueue()
	t.Log(q)
	q.DeQueue()
	t.Log(q)
}

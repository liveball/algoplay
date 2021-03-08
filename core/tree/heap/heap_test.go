package heap

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func randomIntGenerator(count int, max int) []int {
	rand.Seed(time.Now().Unix())
	res := make([]int, count)
	for i := 0; i < count; i++ {
		res[i] = int(rand.Float32() * float32(max))
	}
	return res
}

type Comparator = func(i, j int) bool

func isSorted(list []int) bool {
	var comp Comparator

	comp = func(i, j int) bool {
		return list[i] <= list[j]
	}

	for i := 0; i < len(list)-1; i++ {
		if !comp(i, i+1) {
			return false
		}
	}
	return true
}

func Test_heapSort(t *testing.T) {
	a := []int{2, 1, 4, 3, 5, 7, 6, 9, 8, 10}

	fmt.Println("before sort", a)

	HeapSort(a)

	fmt.Println("after sort", a, isSorted(a))

}

func Test_RemoveMin(t *testing.T) {
	a := []int{2, 1, 4, 3, 5, 7, 6, 9, 8, 10}

	fmt.Println("before build", a)

	BuildMinHeap(a, len(a)-1)

	fmt.Println("after build", a)

	RemoveMin(a, len(a)-1)

	fmt.Println("remove heap top", a)

}

func Test_heap(t *testing.T) {
	cnt := 10
	h := NewHeap(cnt)

	for i := 1; i <= cnt; i++ {
		h.insert(i)
	}

	fmt.Println("before build", h.a)

	h.buildHeap()

	fmt.Println("after build", h.a)

	h.removeMax()
	fmt.Println("after remove", h.a)
}

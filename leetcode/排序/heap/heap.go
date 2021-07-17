package heap

import (
	"fmt"
	"sync"
)

func HeapSortByMinHeapTest() {
	a := []int{2, 1, 3, 4, 6, 5}

	fmt.Println("before sort:", a)

	var wg sync.WaitGroup
	n := len(a)
	wg.Add(n)
	var once sync.Once

	for i := 0; i < n; i++ {
		go func() {
			once.Do(func() {
				HeapSortByMinHeap(a)
				fmt.Println("only once sort")
			})

			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("after sort:", a)
}

func HeapSortByMinHeap(a []int) {
	heapSize := len(a)

	makeMinHeap(a, heapSize)

	for i := len(a) - 1; i > 0; i-- {
		swap(a, 0, i)
		heapSize--

		minHeapify(a, 0, heapSize)
	}
}

func makeMinHeap(a []int, n int) {
	for i := n / 2; i >= 0; i-- {
		minHeapify(a, i, n)
	}
}

func minHeapify(a []int, i, n int) {
	minPos := i
	for {
		if 2*i < n && a[minPos] < a[2*i] {
			minPos = 2 * i
		}

		if 2*i+1 < n && a[minPos] < a[2*i+1] {
			minPos = 2*i + 1
		}

		if minPos == i {
			break
		}

		swap(a, i, minPos)
		i = minPos
	}
}

func swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}

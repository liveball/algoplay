package heap

import "fmt"

type Heap struct {
	a     []int //数组，从下标1开始存储数据
	n     int   //堆可以存储的最大数据个数
	count int   //堆中已经存储的数据个数
}

func NewHeap(capacity int) Heap {
	return Heap{
		a:     make([]int, capacity+1),
		n:     capacity,
		count: 0,
	}
}

//top-max heap -> heapify from down to up
func (heap *Heap) insert(data int) {
	//defensive
	if heap.count == heap.n {
		return
	}

	heap.count++
	heap.a[heap.count] = data

	//compare with parent node
	i := heap.count
	parent := i / 2
	for parent > 0 && heap.a[parent] < heap.a[i] {
		swap(heap.a, parent, i)
		i = parent
		parent = i / 2
	}
}

func (heap *Heap) buildHeap() {
	for i := heap.count / 2; i >= 1; i-- {
		maxHeapify(heap.a, i, heap.count)
	}
}

//heapfify from up to down
func (heap *Heap) removeMax() {

	//defensive
	if heap.count == 0 {
		return
	}

	//swap max and last
	swap(heap.a, 1, heap.count)
	heap.count--

	//heapify from up to down
	heapifyUpToDown(heap.a, heap.count)
}

//heapify
func heapifyUpToDown(a []int, count int) {

	for i := 1; i <= count/2; {

		maxIndex := i
		if a[i] < a[i*2] {
			maxIndex = i * 2
		}

		if i*2+1 <= count && a[maxIndex] < a[i*2+1] {
			maxIndex = i*2 + 1
		}

		if maxIndex == i {
			break
		}

		swap(a, i, maxIndex)
		i = maxIndex
	}

}

func swap(a []int, i, j int) {
	t := a[i]
	a[i] = a[j]
	a[j] = t
}

func HeapSort(a []int) {
	heapSize := len(a)

	BuildMinHeap(a, heapSize)
	// BuildMaxHeap(a, heapSize)

	fmt.Println("after build", a)

	for i := len(a) - 1; i > 0; i-- {
		swap(a, 0, i) //交换 堆顶元素与最后一个元素
		heapSize--

		minHeapify(a, 0, heapSize) //堆化 小顶堆 i-1 每次从堆顶开始堆化，且元素个数减1
		// maxHeapify(a, 0, heapSize) //堆化 大顶堆

		fmt.Println(a)
	}
}

func BuildMinHeap(a []int, n int) {
	for i := n / 2; i >= 0; i-- {
		minHeapify(a, i, n) //每次从i节点往下开始堆化，且元素个数不变
	}
}

func BuildMaxHeap(a []int, n int) {
	for i := n / 2; i >= 0; i-- {
		maxHeapify(a, i, n)
	}
}

func RemoveMin(a []int, n int) {
	if n == 0 {
		return
	}

	swap(a, 0, n)
	n-- //删除一个元素
	minHeapify(a, 0, n)
}

func minHeapify(a []int, i int, count int) { //a:数组，i:父节点位置，count:堆元素个数
	var minPos int

	for {
		minPos = i

		if 2*i < count && a[minPos] > a[2*i] {
			minPos = 2 * i
		}

		if 2*i+1 < count && a[minPos] > a[2*i+1] {
			minPos = 2*i + 1
		}

		if minPos == i {
			break
		}

		swap(a, i, minPos)
		i = minPos
	}
}

func maxHeapify(a []int, i int, count int) {
	var maxPos int

	for {
		maxPos = i

		if 2*i < count && a[maxPos] < a[2*i] {
			maxPos = 2 * i
		}

		if 2*i+1 < count && a[maxPos] < a[2*i+1] {
			maxPos = 2*i + 1
		}

		if maxPos == i {
			break
		}

		swap(a, i, maxPos)
		i = maxPos
	}
}

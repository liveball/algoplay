package sort

import (
	"algoplay/common"
)

// HeapSort is heap sort for List
func HeapSort(list common.List, comp common.Comparator) {
	BuildMaxHeap(list, comp)

	for i := list.Length() - 1; i > 0; i-- {
		exchange(list, 0, i)
		list = list.Slice(0, list.Length()-1)

		maxHeapify(list, 0, comp)
	}
}
// BuildMaxHeap 构建最大堆
func BuildMaxHeap(list common.List, comp common.Comparator) {
	lastRoot := list.Length()/2 - 1
	for i := lastRoot; i > -1; i-- {
		maxHeapify(list, i, comp)
	}
}

// i 的左子节点为 2i+1
func left(i int) int { return 2*i + 1 }

// i 的右子节点为 2i+2
func right(i int) int { return 2*i + 2 }

// 对list的第i个元素代表的子树最大堆化
func maxHeapify(list common.List, i int, comp common.Comparator) {
	for {
		l := left(i)
		r := right(i)

		largest := i // 令largest指向该子树三个节点中最大的那个

		if l < list.Length() && comp(i, l) {
			largest = l
		}

		if r < list.Length() && comp(largest, r) {
			largest = r
		}

		if largest != i {
			exchange(list, i, largest)
			i = largest
		} else {
			break
		}
	}
}



// IsMaxHeap 判断是否符合最大堆性质：list[parent(i)] >= list[i] for all i > 0
func IsMaxHeap(list common.List, comp common.Comparator) bool {
	for i := 1; i < list.Length(); i++ {
		if !comp(i, parent(i)) {
			return false
		}
	}
	return true
}

// i 的父节点为 向下取整[(i-1)/2]
func parent(i int) int {
	if (i-1)%2 == 0 {
		return (i - 1) / 2
	}
	return (i - 2) / 2
}

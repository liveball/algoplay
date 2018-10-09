package sort

import (
	"sync"

	"github.com/algoplay/common"
)

func QuickSort(list common.List, comparator common.Comparator) {
	quickSort(list, 0, list.Length()-1, comparator)
}

func QuickSortConcurrent(list common.List, comparator common.Comparator) {
	wg := &sync.WaitGroup{}
	quickSortConcurrent(list, 0, list.Length()-1, comparator, wg)
	wg.Add(-1)
	wg.Wait()
}

func quickSort(list common.List, low, high int, comparator common.Comparator) {
	if low >= high {
		return
	}
	mid := partition(list, low, high, comparator)
	quickSort(list, low, mid-1, comparator)
	quickSort(list, mid+1, high, comparator)
}

func quickSortConcurrent(list common.List, low, high int, comparator common.Comparator, wg *sync.WaitGroup) {
	wg.Add(1)
	if low >= high {
		return
	}
	mid := partition(list, low, high, comparator)
	go func() {
		quickSortConcurrent(list, low, mid-1, comparator, wg)
		wg.Add(-1)
	}()
	go func() {
		quickSortConcurrent(list, mid+1, high, comparator, wg)
		wg.Add(-1)
	}()
}

func partition(list common.List, low, high int, comparator common.Comparator) (mid int) {
	// 将数组切分为list[low..j-1] <= a[j] <= a[j+1, high]
	i, j := low, high+1 // 左右扫描指针
	// low is 切分元素
	for {
		i++
		for comparator(i, low) {
			if i == high {
				break
			}
			i++
		}
		j--
		for comparator(low, j) {
			if low == j {
				break
			}
			j--
		}
		if i >= j {
			break
		}
		exch(list, i, j)
	}
	exch(list, low, j)
	return j
}

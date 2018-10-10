package sort

import (
	"sync"

	"github.com/liveball/algoplay/common"
)

var pool = common.New(2)

//QuickSort serial quick sort
func QuickSort(list common.List, comparator common.Comparator) {
	quickSort(list, 0, list.Length()-1, comparator)
}

//QuickSortConcurrent concurrent quick sort
func QuickSortConcurrent(list common.List, comparator common.Comparator) {
	wg := &sync.WaitGroup{}
	quickSortConcurrent(list, 0, list.Length()-1, comparator, wg)
	wg.Wait()
}

//Close close pool
func Close() {
	pool.Close() //让工作池停止工作， 等待所有现有的工作完成
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
	if low >= high {
		return
	}
	mid := partition(list, low, high, comparator)

	wg.Add(2)
	// go func() {
	// 	quickSortConcurrent(list, low, mid-1, comparator, wg)
	// 	wg.Done()
	// }()
	// go func() {
	// 	quickSortConcurrent(list, mid+1, high, comparator, wg)
	// 	wg.Done()
	// }()

	go func() {
		pool.Go(func() {
			quickSortConcurrent(list, low, mid-1, comparator, wg)
		})
		wg.Done()
	}()
	go func() {
		pool.Go(func() {
			quickSortConcurrent(list, mid+1, high, comparator, wg)
		})
		wg.Done()
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
		exchange(list, i, j)
	}
	exchange(list, low, j)
	return j
}

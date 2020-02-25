package main

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func Test_binarySearch(t *testing.T) {
	//c := func(i int) bool { return data[i] > 9 }
	//i := sort.Search(len(data)-1, c)
	//println(i)

	m := binNormalSearch([]int{1, 3, 5, 7, 8}, 3)
	fmt.Println(m)

	n := BinarySearch([]int{2, 4, 5, 6, 9}, 5)
	fmt.Println("BinarySearch", n)

	sf := BinarySearchFirst([]int{2, 3, 5, 3, 9}, 3)
	fmt.Println("BinarySearchFirst", sf)

	sl := BinarySearchLast([]int{0, 1, 1, 1, 1}, 1)
	fmt.Println("BinarySearchLast", sl)

	sfg := BinarySearchFirstGT([]int{2, 1, 5, 3, 9}, 3)
	fmt.Println("BinarySearchFirstGT", sfg)

	slt := BinarySearchLastLT([]int{5, 1, 2, 3, 9}, 4)
	fmt.Println("BinarySearchLastLT", slt)

	//在数组 [2, 4, 5, 6, 9] 中查找数字3，就会返回数字4的位置；在数组 [0, 1, 1, 1, 1] 中查找数字1，就会返回第一个数字1的位置

	//比如在数组 [2, 4, 5, 6, 9] 中查找数字3，还是返回数字4的位置，这跟上面那查找方式返回的结果相同，
	//因为数字4在此数组中既是第一个不小于目标值3的数，
	//也是第一个大于目标值3的数，所以 make sense；
	//在数组 [0, 1, 1, 1, 1] 中查找数字1，就会返回坐标5，通过对比返回的坐标和数组的长度，
	//我们就知道是否存在这样一个大于目标值的数

}

// 普通二分查找
func binNormalSearch(data []int, key int) int {

	left, right := 0, len(data)

	f := func(i int) bool { return data[i] >= key }

	for left < right {
		mid := (left + right) >> 1

		if !f(mid) {
			left = mid + 1 //preserves (维护) f(left-1) == false
		} else {
			right = mid // preserves f(right) == true
		}

	}

	// left == right, f(left-1) == false, and f(right) (= f(left)) == true  =>  answer is left.
	return left
}

//ref: https://www.cnblogs.com/grandyang/p/6854825.html
//https://github.com/wangzheng0822/algo/blob/master/go/15_binarysearch/binarysearch.go

//第一类 查找和目标值完全相等的数
func BinarySearch(data []int, target int) (mid int) {
	left, right := 0, len(data)
	for left < right {
		mid := left + (right-left)/2 //(left+right) >> 1
		//fmt.Println("mid", mid, "mid2", (left+right) >> 1)
		if data[mid] == target {
			return mid
		} else if data[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return -1
}

//查找第一个等于给定值的元素
func BinarySearchFirst(data []int, target int) int {
	left, right := 0, len(data)

	for left < right {
		mid := (left + right) >> 1

		if data[mid] < target {
			left = mid + 1
		} else if data[mid] > target {
			right = mid
		} else {
			if mid == 0 || data[mid-1] != target {
				return mid
			} else {
				right = mid
			}
		}
	}

	return -1
}

//查找最后一个值等于给定值的元素
func BinarySearchLast(data []int, target int) int {
	left, right := 0, len(data)

	for left < right {
		mid := (left + right) >> 1

		if data[mid] < target {
			left = mid + 1
		} else if data[mid] > target {
			right = mid
		} else {
			if mid == len(data)-1 || data[mid+1] != target {
				return mid
			} else {
				left = mid + 1
			}
		}
	}

	return -1
}

//查找第一个大于等于给定值的元素
func BinarySearchFirstGT(data []int, target int) int {
	left, right := 0, len(data)

	for left < right {
		mid := (left + right) >> 1

		if data[mid] < target {
			left = mid + 1
		} else if data[mid] > target {
			right = mid
		} else {
			if mid != len(data)-1 || data[mid+1] > target {
				return mid + 1
			} else {
				left = mid + 1
			}
		}
	}

	return -1
}

//查找最后一个小于等于给定值的元素
func BinarySearchLastLT(data []int, target int) int {
	left, right := 0, len(data)

	for left < right {
		mid := (left + right) >> 1

		if data[mid] < target {
			left = mid + 1
		} else if data[mid] > target {
			right = mid
		} else {
			if mid == 0 || data[mid-1] < target {
				return mid
			} else {
				right = mid
			}
		}
	}

	return -1
}

var mids = "268104,389088,1724598,4931952,8820267"

func TestFindArr(t *testing.T) {
	mid := "8820267"
	midArr := strings.Split(mids, ",")
	now := time.Now()
	// res := orderSearch(midArr, mid)
	res := binarySearch(midArr, mid)
	println(res)
	elapsed := time.Since(now)
	fmt.Println(elapsed)
}

func orderSearch(s []string, e string) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func binarySearch(s []string, e string) bool {
	var low, mid, high int
	high = len(s) - 1
	for low <= high {
		mid = (low + high) / 2
		if s[mid] == e {
			return true
		}
		if s[mid] < e {
			low = mid + 1
		}
		if s[mid] > e {
			high = mid - 1
		}
		println("mid", mid, s[mid], "low", low, "high", high)
	}
	return false
}

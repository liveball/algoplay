package main

import "fmt"

var (
	tests = map[int][]int{
		3: []int{2, 4, 5, 6, 9},
		1: []int{0, 1, 1, 1, 1},
	}
)

func main() {
	for k, v := range tests {
		res := BinarySearchFirstGreaterTarget(v, k)
		fmt.Println(k, v, res)
	}

}

//查找第一个大于等于给定值的元素
func BinarySearchFirstGreaterTarget(a []int, target int) int {
	n := len(a)
	if n == 0 {
		return -1
	}

	low := 0
	high := n - 1

	for low <= high {
		mid := (low + high) >> 1

		if a[mid] > target {
			high = mid - 1

		} else if a[mid] < target {
			low = mid + 1

		} else {
			if mid != n-1 && a[mid+1] > target {
				return mid + 1
			} else {
				low = mid + 1
			}
		}
	}

	return -1
}

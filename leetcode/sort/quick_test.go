package sort

import (
	"sort"
	"testing"
)

var (
	nums = []int{
		3, 2, 5, 1, 6,
	}
)

func Test_quickSort(t *testing.T) {
	sort.Ints(nums)
	//QuickSort(nums)
	t.Log(nums)
}

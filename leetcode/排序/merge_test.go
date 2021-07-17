package test

import (
	"testing"
)

func Test_mergeSort(t *testing.T) {
	//sort.Ints(nums)
	mergeSort(nums)
	t.Log(nums)
}

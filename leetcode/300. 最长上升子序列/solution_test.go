package _00_最长上升子序列

import (
	"fmt"
	"testing"
)

var (
	nums = []int{10, 9, 2, 5, 3, 7, 101, 18}
)

func Test_lengthOfLIS(t *testing.T) {
	ret := lengthOfLIS_Recursion(nums)

	fmt.Println(ret)
}

func Test_lengthOfLIS2(t *testing.T) {
	ret := dp_BottomUp(nums)

	fmt.Println(ret)
}

func Test_lengthOfLIS3(t *testing.T) {
	ret := lengthOfLIS_BinarySearch(nums)

	fmt.Println(ret)
}

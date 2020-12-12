package queue

import (
	"fmt"
	"testing"
)

var (
	nums = []int{1, 3, -1, -3, 5, 3, 6, 7}
)

func Test_maxSlidingWindow(t *testing.T) {
	fmt.Println(maxSlidingWindow(nums, 3))
}

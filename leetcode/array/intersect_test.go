package array

import (
	"fmt"
	"testing"
)

func Test_intersect1(t *testing.T) {
	fmt.Println(intersect1([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersect1([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
	fmt.Println(intersect1([]int{1, 2, 3, 4, 4, 13}, []int{1, 2, 3, 9, 10}))
}

func Test_intersect2(t *testing.T) {
	fmt.Println(intersect2([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersect2([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
	fmt.Println(intersect2([]int{1, 2, 3, 4, 4, 13}, []int{1, 2, 3, 9, 10}))
}

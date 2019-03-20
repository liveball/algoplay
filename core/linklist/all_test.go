package main

import (
	"fmt"
	"testing"
)

func Test_reverseList(t *testing.T) {
	in := []int{1, 2, 3, 4, 5}
	fmt.Printf("【input】:%v 【output】:%v\n", in, l2s(reverseList(s2l(in))))
}

func Test_swapList(t *testing.T) {
	in := []int{1, 2, 3, 4, 5}
	fmt.Printf("【input】:%v 【output】:%v\n", in, l2s(swapList(s2l(in))))
}

func Test_hasCycle(t *testing.T) {
	in := []int{1, 2, 3, 4, 5}
	println(hasCycle(s2l(in)))
}

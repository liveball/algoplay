package main

import (
	"fmt"
	"testing"
)

func Test_CycleLinkedList(t *testing.T) {
	in := []int{1, 2, 3, 4, 5}

	cy := CycleLinkedList(s2l(in))
	fmt.Println(HasCycle(cy))
}

func Test_HasCycle(t *testing.T) {
	in := []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	fmt.Println(HasCycle(s2l(in)))
}

func Test_ReverseList(t *testing.T) {
	in := []int{1, 2, 3, 4, 5}

	Print(s2l(in))
	fmt.Printf("【input】:%v 【output】:%v\n", in, l2s(ReverseList(s2l(in))))
}

func Test_InsertAfter(t *testing.T) {
	in := []int{1, 2, 3, 5}

	head := s2l(in)

	n3 := head.Next.Next
	Print(head)
	Print(InsertAfter(n3, 4))
}

func Test_InsertToHead(t *testing.T) {
	in := []int{1, 3, 4, 5}

	head := s2l(in)

	Print(head)
	Print(InsertToHead(head, 2))
}

func Test_InsertToTail(t *testing.T) {
	in := []int{1, 2, 3, 4}

	head := s2l(in)

	Print(head)
	Print(InsertToTail(head, 5))
}

func Test_InsertBefore(t *testing.T) {
	in := []int{1, 3, 4, 5}

	head := s2l(in)

	n2 := head.Next
	Print(head)
	Print(InsertBefore(head, n2, 2))
}

func Test_DelNodeByVal(t *testing.T) {
	in := []int{1, 2, 3, 4, 5}
	fmt.Printf("【input】:%v 【output】:%v\n", in, l2s(DelNodeByVal(s2l(in), 1)))
}

func Test_DelNode(t *testing.T) {
	in := []int{1, 2, 3, 4, 5}

	head := s2l(in)

	n := head.Next.Next

	Print(head)
	Print(DelNode(head, n))
}

func Test_DelBottomN(t *testing.T) {
	in := []int{1, 2, 3, 4, 5}

	//Print(s2l(in))

	Print(DelBottomN(s2l(in), 3))
}

func Test_FindByIndex(t *testing.T) {
	in := []int{1, 2, 3, 4, 5}

	head := s2l(in)

	Print(head)
	Print(FindByIndex(head, 2))
}

func Test_SwapList(t *testing.T) {
	in := []int{1, 2, 3, 4, 5}
	fmt.Printf("【input】:%v 【output】:%v\n", in, l2s(SwapList(s2l(in))))
}

func Test_MergeSortedList(t *testing.T) {
	l1 := s2l([]int{1, 3, 5, 7, 9})
	l2 := s2l([]int{2, 4, 6, 8, 10})

	Print(l1)
	Print(l2)

	l := MergeSortedList(l1, l2)
	Print(l)
}


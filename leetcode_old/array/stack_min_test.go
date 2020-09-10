package main

import (
	"fmt"
	"testing"
)

func Test_stackMin(t *testing.T) {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)

	minStack.PrintElements()
	minStack.PrintMin()
	fmt.Println(minStack.GetMin()) //--> 返回 -3.

	minStack.Pop()

	minStack.PrintElements()
	minStack.PrintMin()

	fmt.Println(minStack.Top())    //--> 返回 0.
	fmt.Println(minStack.GetMin()) //--> 返回 -2.
}

package array

import (
	"fmt"
	"testing"
)

func Test_stackMax(t *testing.T) {
	maxStack := ConstructorMaxStack()

	maxStack.Push(1)

	p1 := maxStack.Pop()

	p2 := maxStack.Top()
	p3 := maxStack.PeekMax()
	p4 := maxStack.PopMax()

	maxStack.PrintMaxStack()

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
	fmt.Println(p4)
}

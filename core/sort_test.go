package core

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	//test int type
	ints := &In{}
	for i := 0; i < 10; i++ {
		ints.Add(i)
	}

	fmt.Println(ints.Get()) //before sort
	BubbleSort(ints.Get())
	fmt.Println(ints.Get()) //after sort

	//test string type
	strs := &In{}
	for i := 'a'; i < 'z'; i++ {
		strs.Add(fmt.Sprintf("%c ", i))
	}

	fmt.Println(strs.Get()) //before sort
	BubbleSort(strs.Get())
	fmt.Println(strs.Get()) //after sort

}

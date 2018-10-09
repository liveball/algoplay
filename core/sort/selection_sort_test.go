package sort

import (
	"algoplay/core"
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	ints := &core.In{}
	for i := 0; i < 10; i++ {
		ints.Add(i)
	}
	fmt.Println(ints.Get())
	//
	//fmt.Println(ints.Get()) //before sort
	SelectionSort(ints.Get())
	//fmt.Println(ints.Get())
}

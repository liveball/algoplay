package sort

import (
	"fmt"
	"testing"

	"github.com/liveball/algoplay/tools"
)

func TestSelectionSort(t *testing.T) {
	ints := &tools.In{}
	for i := 0; i < 10; i++ {
		ints.Add(i)
	}
	fmt.Println(ints.Get())
	//
	//fmt.Println(ints.Get()) //before sort
	SelectionSort(ints.Get())
	//fmt.Println(ints.Get())
}

package old_sort

import (
	"fmt"
	"testing"
	"time"

	"algoplay/tools"
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

func Test_selectionSort(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1, 9, 6}
	start := time.Now()
	selection(arr)
	elapsed := time.Since(start)
	fmt.Println(elapsed)
	fmt.Println(arr)
}

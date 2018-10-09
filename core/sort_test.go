package core

import (
	"fmt"
	"testing"

	"github.com/algoplay/tools"
)

func TestBubbleSort(t *testing.T) {
	//test int type
	ints := &In{}
	for _, v := range tools.New().Numbers(0, 99, 30) {
		ints.Add(v)
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

func TestSelectionSort(t *testing.T) {
	//test int type
	ints := &In{}
	for _, v := range tools.New().Numbers(0, 99, 30) {
		ints.Add(v)
	}
	fmt.Println(ints.Get()) //before sort
	SelectionSort(ints.Get())
	fmt.Println(ints.Get()) //after sort

	//test string type
	strs := &In{}
	for _, v := range tools.New().Strings() {
		strs.Add(v)
	}

	fmt.Println(strs.Get()) //before sort
	SelectionSort(strs.Get())
	fmt.Println(strs.Get()) //after sort
}

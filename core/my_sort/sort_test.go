package my_sort

import (
	"fmt"
	"testing"
)

func Test_MergeSort(t *testing.T) {
	a := []int{2, 1, 9, 6, 8, 7, 4, 5, 10, 3}

	fmt.Println("排序前：", a)

	MergeSort(a)

	fmt.Println("排序后：", a)
}

func Test_QuickSort(t *testing.T) {
	a := []int{2, 1, 9, 6, 8, 7, 4, 5, 10, 3}

	fmt.Println("排序前：", a)

	QuickSort(a)

	fmt.Println("排序后：", a)
}

func Test_BubbleSort(t *testing.T) {
	a := []int{2, 1, 9, 6, 8, 7, 4, 5, 10, 3}

	fmt.Println("排序前：", a)

	BubbleSort(a)

	fmt.Println("排序后：", a)
}

func Test_InsertionSort(t *testing.T) {
	a := []int{2, 1, 9, 6, 8, 7, 4, 5, 10, 3}

	fmt.Println("排序前：", a)

	InsertionSort(a)

	fmt.Println("排序后：", a)
}

func Test_SelectionSort(t *testing.T) {
	a := []int{2, 1, 9, 6, 8, 7, 4, 5, 10, 3}

	fmt.Println("排序前：", a)

	SelectionSort(a)

	fmt.Println("排序后：", a)
}
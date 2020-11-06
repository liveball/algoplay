package main

import "fmt"

func main() {
	arr := []int{3, 1, 2, 4}

	selectSort(arr)

	fmt.Println(arr)
}

func selectSort(arr []int) {
	if len(arr) == 0 {
		return
	}

	for i := 0; i < len(arr)-1; i++ {
		min := i

		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}

		arr[min], arr[i] = arr[i], arr[min]
	}
}

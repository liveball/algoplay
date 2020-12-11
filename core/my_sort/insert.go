package main

import "fmt"

func insertSort(arr []int) {
	if len(arr) == 0 {
		return
	}

	for i := 1; i < len(arr)-1; i++ {
		temp := arr[i]
		j := i - 1
		for ; j >= 0;j--{
			if arr[j] > temp{
				arr[j+1] = arr[j]
			}else {
				break
			}
		}

		arr[j+1] = temp
	}
}

func main() {
	arr := []int{3, 1, 2, 4}

	insertSort(arr)

	fmt.Println(arr)
}

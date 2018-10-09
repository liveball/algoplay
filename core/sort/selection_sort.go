package sort

import (
	"fmt"
)

func SelectionSort(data interface{}) interface{} {
	arr := data.([]interface{})
	length := len(arr)
	minIndex := 0
	tmp := 0
	for i_k, _ := range arr {
		switch arr[i_k].(type) {
		case int:
			minIndex = i_k

			for j := i_k + 1; j < length; j++ {
				m := arr[j].(int)
				n := arr[minIndex].(int)
				if m > n {
					minIndex = j
				}
			}
			tmp = arr[i_k].(int)
			arr[i_k] = arr[minIndex]
			arr[minIndex] = tmp
		}

	}
	fmt.Println(arr)

	return arr
}

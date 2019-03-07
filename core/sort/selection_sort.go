package sort

import (
	"fmt"
)

func SelectionSort(data interface{}) interface{} {
	arr := data.([]interface{})
	length := len(arr)

	for i := range arr {
		minIndex := 0
		tmp := 0
		switch arr[i].(type) {
		case int:
			minIndex = i

			for j := i + 1; j < length; j++ {
				m := arr[j].(int)
				n := arr[minIndex].(int)
				if m > n {
					minIndex = j
				}
			}
			tmp = arr[i].(int)
			arr[i] = arr[minIndex]
			arr[minIndex] = tmp
		}

	}
	fmt.Println(arr)

	return arr
}

package sort

import (
	"fmt"
)

func selection(arr []int) {
	l := len(arr)
	if l <= 1 {
		return
	}
	for i := 0; i < l; i++ { //有序区域 R[0..i]
		min := i
		for j := i + 1; j < l-1; j++ { //无序区域 R[i+1..n-1] 查找
			if arr[min] > arr[j] { //选出最小的
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i] //有序区域增加最小元素
	}
}

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

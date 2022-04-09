package main

import (
	"fmt"
)

func qsortTest() {
	arr := []int{3, 1, 2, 4}

	qsort(arr, 0, len(arr)-1)

	fmt.Println(arr)
}

func qsort(arr []int, l, r int) {
	if l >= r {
		return
	}

	p := partition2(arr, l, r)
	qsort(arr, l, p-1)
	qsort(arr, p+1, r)
}

func partition2(arr []int, l, r int) int {
	i := l      //标记分割位置
	v := arr[l] //记录比较值
	//从分割位置右边的第一个元素开始比较
	for j := l + 1; j <= r; j++ { //增加右边区间
		if arr[j] < v {
			i++ //增加左边区间
			swap(arr, i, j)
		}
	}

	swap(arr, l, i) //交换分割位置
	return i
}

func swap(arr []int, i, j int) {
	arr[j], arr[i] = arr[i], arr[j]
}

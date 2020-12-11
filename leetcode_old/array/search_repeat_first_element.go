package main

import "fmt"

func searchFirstRepeatElement(a []int) (res int) {

	for i := 0; i < len(a)-1; i++ {
		if a[i]^a[i+1] == 0 {
			res = i
		}
	}

	return
}

func main() {
	a := [][]int{
		{1, 1, 2, 3, 4, 5},
		{1, 2, 2, 3, 4, 5},
		{1, 2, 3, 3, 4, 5},
		{1, 2, 3, 4, 4, 5},
		{1, 2, 3, 4, 5, 5},
	}
	for _, v := range a {
		fmt.Println(searchFirstRepeatElement(v))
	}
}

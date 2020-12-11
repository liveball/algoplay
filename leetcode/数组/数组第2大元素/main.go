package main

import (
	"fmt"
)

func main() {
	fmt.Println(getSecMax([]int{1, 2, 3, 4, 5, 6}))
}

func getSecMax(nums []int) int {
	max := 0
	secMax := 0
	for _, num := range nums {
		if num > max {
			secMax = max
			max = num
			continue
		}
		if num > secMax {
			secMax = num
		}
	}

	return secMax
}

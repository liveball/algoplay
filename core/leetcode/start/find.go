package main

import "fmt"

func main() {
	a := []int{2, 2, 1}
	fmt.Println(a)
        //singleNumber(a)
	println(singleNumber(a))
}

func singleNumber(nums []int) int {
	for k, v := range nums {
		var j int
		for _, vv := range nums {
			if v == vv {
				j++
				println("j:",j)
			}
                        if j==1{
                          break
                        }
		}
		if j == 1 {
			return k
		}
	}
	return 0
}

package main

import (
	"fmt"
)

// go build -o main -gcflags "-N -l" && GODEBUG=gctrace=1   ./main

func main() {
	// a := []int{2, 7, 11, 15}
	// fmt.Println(twoSum(a, 9))

	var vs []int
	fmt.Println(vs)
	vs = make([]int, 0, 2)
	vs = append(vs, 10)
	inst(vs)
	fmt.Println(vs)

	fmt.Println(111)
}

func inst(vals []int) {
	vals[0] = 20
	// vals = append(vals, 20)
}

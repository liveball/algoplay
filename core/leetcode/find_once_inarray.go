package main

import "fmt"

func main() {
	a := []int{2, 2, 1}
	fmt.Println(a)
	//singleNumber(a)
	println(singleNumber(a))
}

// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
// 说明：
// 你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
// 示例 1:
// 输入: [2,2,1]
// 输出: 1
// 示例 2:
// 输入: [4,1,2,1,2]
// 输出: 4
func singleNumber(nums []int) int {
	for k, v := range nums {
		var j int
		for _, vv := range nums {
			if v == vv {
				j++
				println("j:", j)
			}
			if j == 1 {
				break
			}
		}
		if j == 1 {
			return k
		}
	}
	return 0
}

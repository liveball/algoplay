package arr_sort

import (
	"fmt"
	"testing"
)

func Test_threeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}

	fmt.Println(threeSum(nums))
}

//三数之和

//给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。
//
//注意：答案中不可以包含重复的三元组。
//
//例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//
//满足要求的三元组集合为：
//[
//  [-1, 0, 1],
//  [-1, -1, 2]
//]

// a+b = 0-c

func threeSum(nums []int) [][]int {
	var res [][]int
	var res1 []int
	target := 0

	dict := make(map[int]struct{})
	dict1 := make(map[int]struct{})
	for _, v := range nums {
		target2 := target - v

		if _, ok := dict[v]; ok {
			res1 = append(res1, v)
		}
		dict[v] = struct{}{}

		for _, v1 := range nums {
			v2 := target2 - v1
			if _, ok := dict1[v2]; ok {
				res1 = append(res1, v1)
				res1 = append(res1, v2)
			}
			dict1[v1] = struct{}{}
		}

		res = append(res, res1)
	}
	return res
}

package main

import (
	"fmt"
	"reflect"
	"sort"
	"unsafe"
)

// 只出现一次的数字
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
	exist := make(map[int]int)
	for _, v := range nums {
		exist[v] = 0
	}

	for _, v := range nums {
		exist[v]++
	}

	fmt.Println(exist)
	for k, v := range exist {
		if v == 1 {
			return k
		}
	}
	return 0
}

// 求众数
// 给定一个大小为 n 的数组，找到其中的众数。众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
// 你可以假设数组是非空的，并且给定的数组总是存在众数。
// 示例 1:
// 输入: [3,2,3]
// 输出: 3
// 示例 2:

// 输入: [2,2,1,1,1,2,2]
// 输出: 2
func majorityElement(nums []int) int {
	exist := make(map[int]int)
	for _, v := range nums {
		exist[v] = 0
	}

	for _, v := range nums {
		exist[v]++
	}

	fmt.Println(exist)
	for k, v := range exist {
		if v >= len(nums)/2 {
			return k
		}
	}
	return 0
}

// 合并两个有序数组
// 给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。

// 说明:

// 初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
// 你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
// 示例:

// 输入:
// nums1 = [1,2,3,0,0,0], m = 3
// nums2 = [2,5,6],       n = 3

// 输出: [1,2,2,3,5,6]
func merge(nums1 []int, m int, nums2 []int, n int) {
	// i := m - 1
	// j := n - 1
	// k := m + n - 1
	// for i >= 0 && j >= 0 {
	// 	if nums1[i] > nums2[j] {
	// 		nums1[k] = nums1[i]
	// 		i--
	// 	} else {
	// 		nums1[k] = nums2[j]
	// 		j--
	// 	}
	// 	k--
	// }

	// for j >= 0 {
	// 	nums1[k] = nums2[j]
	// 	k--
	// 	j--
	// }

	// fmt.Printf("111%#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&nums1)))
	for k, v := range nums1 {
		if v == 0 {
			nums1 = nums1[:k]
			break
		}
	}

	if len(nums1) == cap(nums1) {
		(*reflect.SliceHeader)(unsafe.Pointer(&nums1)).Cap = cap(nums1) + n
	}
	nums1 = append(nums1[:], nums2...)

	fmt.Printf("222%#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&nums1)))

	for i := 0; i < len(nums1)-1; i++ {
		for j := 0; j < len(nums1)-i-1; j++ {
			if nums1[j] > nums1[j+1] {
				t := nums1[j]
				nums1[j] = nums1[j+1]
				nums1[j+1] = t
			}
		}
	}
	sort.Ints(nums1)
	fmt.Println(nums1)
}

// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
// 示例:
// 给定 nums = [2, 7, 11, 15], target = 9
// 因为 nums[0] + nums[1] = 2 + 7 = 9
// 所以返回 [0, 1]
func twoSum(nums []int, target int) []int {
	a := make(map[int]int)

	for k, v := range nums {
		comlement := target - v
		if vv, ok := a[comlement]; ok {
			return []int{vv, k}
		}
		a[v] = k
	}
	return []int{}
}

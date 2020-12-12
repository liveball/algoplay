package main

import (
	"fmt"
	"reflect"
	"sort"
	"unsafe"
)

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

	fmt.Printf("222 %#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&nums1)))

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

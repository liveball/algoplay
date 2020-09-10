package main

import (
	"sort"
)

func intersect1(nums1, nums2 []int) []int {
	m0 := make(map[int]int)

	for _, v := range nums1 {
		m0[v]++
	}

	//fmt.Println(m0)
	nums3 := make([]int, len(m0))
	k := 0
	for _, v := range nums2 {
		if m0[v] > 0 {
			m0[v]--
			nums3[k] = v
			k++
		}
	}

	//fmt.Println(m0, nums3, k)

	return nums3[0:k]
}

func intersect2(nums1, nums2 []int) []int {
	var i, j, k int
	sort.Ints(nums1)
	sort.Ints(nums2)

	cnt := 0
	if len(nums1) > len(nums2) {
		cnt = len(nums1)
	} else {
		cnt = len(nums2)
	}

	nums3 := make([]int, cnt)
	for i < len(nums1) && j < len(nums2) { //反复以下步骤，直到任意一个数组终止
		if nums1[i] < nums2[j] { //如果两个指针的元素不相等，我们将小的一个指针前移。
			i++
		} else if nums1[i] > nums2[j] { //如果两个指针的元素不相等，我们将小的一个指针前移。
			j++
		} else { //设定两个为0的指针，比较两个指针的元素是否相等。如果指针的元素相等，我们将两个指针一起向前移动，并且将相等的元素放入空白数组。
			nums3[k] = nums1[i]
			i++
			j++
			k++
		}
	}

	return nums3[:k]
}

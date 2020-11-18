package main

import "fmt"

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))          //5
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4)) //4
}

// func findKthLargest(nums []int, k int) int {
//     sort.Ints(nums)
//     fmt.Println(nums)
//     return nums[len(nums)-k]
// }

func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, l, r, k int) int {
	p := partition(nums, l, r)
	if p == k {
		return nums[p]
	} else if p < k { //在分区点右边
		return quickSelect(nums, p+1, r, k)
	} else { //在分区点左边
		return quickSelect(nums, l, p-1, k)
	}
}

func partition(nums []int, l, r int) int {
	pivot := nums[r] //假设最右边的元素的索引为分区点
	i := l           //记录最左边索引
	for j := l; j < r; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++ //移动
		}
	}

	nums[i], nums[r] = nums[r], nums[i]
	return i
}

// 215. 数组中的第K个最大元素
// 在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

// 示例 1:

// 输入: [3,2,1,5,6,4] 和 k = 2
// 输出: 5
// 示例 2:

// 输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
// 输出: 4
// 说明:

// 你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。

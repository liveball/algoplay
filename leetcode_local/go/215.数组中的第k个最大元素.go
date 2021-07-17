/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */

// @lc code=start
func findKthLargest(nums []int, k int) int {
	//直接排序根据数组性质
	// return findKthLargestBySort(nums, k)

	//根据快排性质
	return findKthLargestByQuickSort(nums, k)
}

func findKthLargestBySort(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

func findKthLargestByQuickSort(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}

	return selectLarge(nums, 0, len(nums)-1, len(nums)-k)
}

func selectLarge(nums []int, l, r, k int) int {
	if l == r {
		return nums[l]
	}

	p := partition(nums, l, r)
	if k == p {
		return nums[p]
	} else if k < p {
		return selectLarge(nums, l, p-1, k)
	} else {
		return selectLarge(nums, p+1, r, k)
	}
}

func partition(nums []int, lo, hi int) int {
	pivot := nums[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		if nums[j] < pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	nums[i+1], nums[hi] = nums[hi], nums[i+1]

	return i + 1
}

// @lc code=end


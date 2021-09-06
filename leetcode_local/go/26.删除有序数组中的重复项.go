/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	return removeDuplicates1(nums)
	// return removeDuplicates2(nums)
}

func removeDuplicates1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	finder, last := 0, 0
	for last < len(nums) {
		for nums[finder] == nums[last] {
			finder++
			if finder == len(nums) {
				return last + 1
			}
		}

		nums[last+1] = nums[finder]
		last++
	}

	return last + 1
}

func removeDuplicates2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	lastNum := nums[len(nums)-1]
	i := 0
	for i = 0; i < len(nums)-1; i++ {
		if nums[i] == lastNum {
			break
		}

		if nums[i+1] == nums[i] {
			removeDuplicatesByV(nums, i+1, nums[i])
		}
	}

	return i + 1
}

func removeDuplicatesByV(nums []int, start, v int) int {
	if len(nums) == 0 {
		return 0
	}

	j := start
	for i := start; i < len(nums); i++ {
		if nums[i] != v {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
				j++
			} else {
				j++
			}
		}
	}

	return j
}

// @lc code=end


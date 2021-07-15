/*
 * @lc app=leetcode.cn id=154 lang=golang
 *
 * [154] 寻找旋转排序数组中的最小值 II
 */

// @lc code=start
func findMin(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		if nums[low] < nums[high] {
			return nums[low]
		}

		mid := low + (high-low)>>1
		if nums[low] > nums[high] {
			low = mid + 1
		} else if nums[low] == nums[high] {
			low++
		} else {
			high = mid
		}
	}

	return nums[low]
}

// @lc code=end


/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 */

// @lc code=start
func findMin(nums []int) int {
	low, high := 0, len(nums)-1

	for low < high {
		if nums[low] < nums[high] {
			return nums[low]
		}

		mid := low + (high-low)>>1
		if nums[mid] >= nums[low] {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return nums[low]
}

// @lc code=end


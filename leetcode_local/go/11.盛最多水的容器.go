/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	if len(height) == 0 {
		return 0
	}

	left, right := 0, len(height)-1
	res := 0
	for left <= right {
		area := min(height[left], height[right]) * (right - left)
		res = max(res, area)
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// @lc code=end


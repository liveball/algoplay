/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	left, right := 0, len(height)-1
	lmax, rmax := height[0], height[len(height)-1]
	res := 0
	for left <= right {
		lmax = max(lmax, height[left])
		rmax = max(rmax, height[right])

		if lmax < rmax {
			res += lmax - height[left]
			left++
		} else {
			res += rmax - height[right]
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

// @lc code=end


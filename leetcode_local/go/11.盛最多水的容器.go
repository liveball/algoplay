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
	for left < right {
		width:=right - left
		high:=0
		
		if height[left] < height[right] {
			high = height[left]
			left++
		} else {
			high = height[right]
			right--
		}

		area:=width*high
		if area>res{
           res=area
		}
	}

	return res
}

// @lc code=end


/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	return maxArea1(height)
	//    return maxArea2(height)
}

func maxArea1(height []int) int {
	if len(height) == 0 {
		return 0
	}

	left, right := 0, len(height)-1
	res := 0
	for left < right {
		width := right - left
		high := 0

		if height[left] < height[right] {
			high = height[left]
			left++
		} else {
			high = height[right]
			right--
		}

		area := width * high
		if area > res {
			res = area
		}
	}

	return res
}

func maxArea2(height []int) int {
	if len(height) == 0 {
		return 0
	}

	max := 0
	left := 0
	right := len(height) - 1

	for left < right {
		width := right - left
		high := 0
		if height[left] < height[right] {
			high = height[left]
			left++
		} else {
			high = height[right]
			right--
		}

		res := width * high
		if res > max {
			max = res
		}
	}

	return max
}

// @lc code=end


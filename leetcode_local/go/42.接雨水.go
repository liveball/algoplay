package solution

/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	// return violence(height)//暴力求解
	return memoViolence(height) //备忘录记忆
	// return doublePointer(height) //双指针
}

func violence(height []int) int {
	length := len(height)
	if length == 0 {
		return 0
	}

	res := 0
	for i := 1; i < length-1; i++ {
		lmax := 0
		rmax := 0

		for j := i; j >= 0; j-- {
			lmax = max(lmax, height[j])
		}

		for j := i; j < length; j++ {
			rmax = max(rmax, height[j])
		}

		res += min(lmax, rmax) - height[i]
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

func memoViolence(height []int) int {
	length := len(height)
	if length == 0 {
		return 0
	}

	var res int
	lmax := make([]int, length)
	rmax := make([]int, length)

	lmax[0] = height[0]
	rmax[length-1] = height[length-1]

	for i := 1; i < length; i++ {
		lmax[i] = max(lmax[i-1], height[i])
	}

	for i := length - 2; i >= 0; i-- {
		rmax[i] = max(rmax[i+1], height[i])
	}

	for i := 1; i < length-1; i++ {
		res += min(lmax[i], rmax[i]) - height[i]
	}
	return res
}

func doublePointer(height []int) int {
	length := len(height)
	if length == 0 {
		return 0
	}

	res := 0
	left, right := 0, length-1
	lmax := height[0]
	rmax := height[length-1]
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

// @lc code=end

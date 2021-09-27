/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	// return trap1(height) //精简双指针
	// return trap2(height) //繁琐双指针
	// return trap3(height) //动态规划
	return trap4(height) //单调栈
}

func trap1(height []int) int {
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

func trap2(height []int) int {
	sum := 0
	for i := 0; i < len(height); i++ {
		if i == 0 || i == len(height)-1 {
			continue
		}

		lheight := height[i]
		rheight := height[i]
		for l := i - 1; l >= 0; l-- {
			lheight = max(lheight, height[l])
		}
		for r := i + 1; r < len(height); r++ {
			rheight = max(rheight, height[r])
		}

		h := min(lheight, rheight) - height[i]
		sum += h
	}

	return sum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func trap3(height []int) int {
	if len(height) <= 2 {
		return 0
	}
	size := len(height)
	maxLeft := make([]int, size)
	maxRight := make([]int, size)

	maxLeft[0] = height[0]
	for i := 1; i < size; i++ {
		maxLeft[i] = max(maxLeft[i-1], height[i])
	}

	maxRight[size-1] = height[size-1]
	for i := size - 2; i >= 0; i-- {
		maxRight[i] = max(maxRight[i+1], height[i])
	}

	sum := 0
	for i := 0; i < size; i++ {
		count := min(maxLeft[i], maxRight[i]) - height[i]
		if count > 0 {
			sum += count
		}
	}

	return sum
}

func trap4(height []int) int {
	size := len(height)
	if size <= 2 {
		return 0
	}

	stack := make([]int, size)
	sum := 0
	for i := 0; i < size; i++ {
		if height[i] < height[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else if height[i] == height[stack[len(stack)-1]] {
			// 如果当前遍历的元素（柱子）高度等于栈顶元素的高度，
			// 要跟更新栈顶元素，因为遇到相相同高度的柱子，
			// 需要使用最右边的柱子来计算宽度。
			stack = stack[:len(stack)-1] // 其实这一句可以不加，效果是一样的，但处理相同的情况的思路却变了
			stack = append(stack, i)
		} else {
			for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
				mid := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if len(stack) > 0 {
					h := min(height[stack[len(stack)-1]], height[i]) - height[mid]
					w := i - stack[len(stack)-1] - 1 // 注意减一，只求中间宽度
					sum += h * w
				}
			}
			stack = append(stack, i)
		}
	}

	return sum
}

// @lc code=end


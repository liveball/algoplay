/*
 * @lc app=leetcode.cn id=1124 lang=golang
 *
 * [1124] 表现良好的最长时间段
 */

// @lc code=start
func longestWPI(hours []int) int {
	// return longestWPI1(hours)
	return longestWPI2(hours)
}

func longestWPI1(hours []int) int {
	if len(hours) == 0 {
		return 0
	}

	score := make([]int, len(hours))
	for i := 0; i < len(hours); i++ {
		if hours[i] > 8 {
			score[i] = -1
		} else {
			score[i] = 1
		}
	}

	sum := 0
	res := 0
	for i := 0; i < len(hours); i++ {
		for j := i; j < len(hours); j++ {
			sum += score[j]
			if sum < 0 {
				res = max(res, j-i+1)
			}
		}
		sum = 0
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func longestWPI2(hours []int) int {
	ans := 0
	n := len(hours)
	for i := 0; i < n; i++ {
		if hours[i] > 8 {
			hours[i] = 1
		} else {
			hours[i] = -1
		}
	}

	prefixSum := make([]int, n+1) //0..n
	// for i:=1; i<n+1; i++{
	// prefixSum[i] += prefixSum[i-1] + hours[i-1]
	for i := 0; i < n; i++ {
		prefixSum[i+1] = hours[i] + prefixSum[i]
	}

	//stack
	stack := make([]int, 0, len(prefixSum))
	stack = append(stack, 0)
	for i, v := range prefixSum {
		if i == 0 {
			continue
		}

		if v < prefixSum[stack[len(stack)-1]] {
			stack = append(stack, i)
		}
	}

	for i := len(prefixSum) - 1; i >= 0; i-- {
		if len(stack) == 0 {
			break
		}

		for {
			if len(stack) == 0 {
				break
			}
			if prefixSum[i]-prefixSum[stack[len(stack)-1]] <= 0 {
				break
			}

			ans = max(ans, i-stack[len(stack)-1])
			stack = stack[0 : len(stack)-1]
		}
	}

	return ans

}

// @lc code=end


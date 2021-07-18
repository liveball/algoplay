/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	return maxSubArray1(nums)
	// return maxSubArray2(nums)
	// return maxSubArray3(nums)
	// return maxSubArray4(nums)
}

func maxSubArray1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	res := nums[0]
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		res = max(res, dp[i])
	}

	return res
}

func maxSubArray2(nums []int) int {
	res := -1 << 31

	if len(nums) == 0 {
		return res
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		if dp[i-1] < 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}

	}

	for _, v := range dp {
		res = max(res, v)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// 解法一 DP
func maxSubArray3(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp, res := make([]int, len(nums)), nums[0]
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}
		res = max(res, dp[i])
	}
	return res
}

// 解法二 模拟
func maxSubArray4(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	maxSum, res, p := nums[0], 0, 0
	for p < len(nums) {
		res += nums[p]
		if res > maxSum {
			maxSum = res
		}
		if res < 0 {
			res = 0
		}
		p++
	}
	return maxSum
}

// @lc code=end


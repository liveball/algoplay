/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	return test1(nums)
	// return test2(nums)
}

func test1(nums []int) int {
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

func test2(nums []int) int {
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

// @lc code=end


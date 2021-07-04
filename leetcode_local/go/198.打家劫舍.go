/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */

// @lc code=start
func rob(nums []int) int {
	// return robDP(nums)
	return robNums(nums) //节省内存
}

func robNums(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	nums[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		nums[i] = max(nums[i-2]+nums[i], nums[i-1])
	}

	return nums[len(nums)-1]
}

func robDP(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}

	return dp[len(nums)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// @lc code=end


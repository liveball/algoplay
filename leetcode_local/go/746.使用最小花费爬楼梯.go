/*
 * @lc app=leetcode.cn id=746 lang=golang
 *
 * [746] 使用最小花费爬楼梯
 */

// @lc code=start
func minCostClimbingStairs(cost []int) int {
	cnt := len(cost)
	dp := make([]int, cnt+1)
	dp[0] = cost[0]
	dp[1] = cost[1]

	for i := 2; i < cnt; i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}

	return min(dp[cnt-1], dp[cnt-2])
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// @lc code=end


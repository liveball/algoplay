/*
 * @lc app=leetcode.cn id=309 lang=golang
 *
 * [309] 最佳买卖股票时机含冷冻期
 */

// @lc code=start
func maxProfit(prices []int) int {
	return maxProfit1(prices)
}

// 每次卖出之后都要等一天才能继续交易，也就是第 i 天选择买的时候，要从 i - 2 状态转移。
// 列出状态转移方程。

// dp[i][k][0] = Math.max(dp[i - 1][k][0], dp[i - 1][k][1] + prices[i])
// dp[i][k][1] = Math.max(dp[i - 1][k][1], dp[i - 2][k - 1][0] - prices[i])
//             = Math.max(dp[i - 1][k][1], dp[i - 2][k][0] - prices[i])
// k 同样对状态转移已经没有影响了，可以进一步化简方程。

// dp[i][0] = Math.max(dp[i - 1][0], dp[i - 1][1] + prices[i])
// dp[i][1] = Math.max(dp[i - 1][1], dp[i - 2][0] - prices[i])

func maxProfit1(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	// dp[i][0]: 手上持有股票的最大收益
	// dp[i][1]: 手上不持有股票，并且处于冷冻期中的累计最大收益
	// dp[i][2]: 手上不持有股票，并且不在冷冻期中的累计最大收益
	dp := make([][3]int, n)
	dp[0][0] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])
		dp[i][1] = dp[i-1][0] + prices[i]
		dp[i][2] = max(dp[i-1][1], dp[i-1][2])
	}

	return max(dp[n-1][1], dp[n-1][2])
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// @lc code=end


/*
 * @lc app=leetcode.cn id=714 lang=golang
 *
 * [714] 买卖股票的最佳时机含手续费
 */

// @lc code=start
func maxProfit(prices []int, fee int) int {
	return maxProfit1(prices, fee)

}

// 第六题 k 为正无穷但有手续费
// 在第二题的基础上，添加了手续费。

// 每次交易要支付手续费，只要把手续费从利润中减去即可，可以列出如下两种方程。

// 第一种方程：在每次买入股票时扣除手续费

// dp[i][0] = max(dp[i - 1][0], dp[i - 1][1] + prices[i])
// dp[i][1] = max(dp[i - 1][1], dp[i - 1][0] - prices[i] - fee)
// 第二种方程：在每次卖出股票时扣除手续费

// dp[i][0] = max(dp[i - 1][0], dp[i - 1][1] + prices[i] - fee)
// dp[i][1] = max(dp[i - 1][1], dp[i - 1][0] - prices[i])

func maxProfit1(prices []int, fee int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	dp := make([][2]int, n)
	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}

	return dp[n-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// @lc code=end


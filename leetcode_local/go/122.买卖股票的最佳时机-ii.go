/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lc code=start
func maxProfit(prices []int) int {
	// return maxProfit1(prices)
	return maxProfit2(prices)
}

// 第二题，k = +infinity
// 如果 k 为正无穷，那么就可以认为 k 和 k - 1 是一样的。可以这样改写框架：

// dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
// dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])
//             = max(dp[i-1][k][1], dp[i-1][k][0] - prices[i])

// 我们发现数组中的 k 已经不会改变了，也就是说不需要记录 k 这个状态了：
// dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
// dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i])

func maxProfit1(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}

	return dp[n-1][0]
}

func maxProfit2(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	dp_i_0 := 0
	dp_i_1 := MIN
	for i := 0; i < n; i++ {
		dp_i_0 = max(dp_i_0, dp_i_1+prices[i])
		dp_i_1 = max(dp_i_1, dp_i_0-prices[i])
	}

	return dp_i_0
}

const (
	MIN = -1 << 31
)

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// @lc code=end


/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lc code=start
func maxProfit(prices []int) int {
	return getMax(prices)
}

func getMax(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	// k = +infinity
	// dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
	// dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
	// 			= max(dp[i-1][k][1], dp[i-1][k][0]-prices[i])

	// dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
	// dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])

	dp_i_0 := 0
	dp_i_1 := MIN
	for i := 0; i < n; i++ {
		t := dp_i_0
		dp_i_0 = max(dp_i_0, dp_i_1+prices[i])
		dp_i_1 = max(dp_i_1, t-prices[i])
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


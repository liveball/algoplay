/*
 * @lc app=leetcode.cn id=123 lang=golang
 *
 * [123] 买卖股票的最佳时机 III
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

	// k = 2
	// dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
	// dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])

	k := 2
	dp := make([][][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][]int, k+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = make([]int, 2)
		}
	}
	// fmt.Println(dp)

	for i := 0; i < n; i++ {
		for j := k; j >= 1; j-- {
			if i-1 == -1 {
				// base case：
				// dp[-1][k][0] = dp[i][0][0] = 0
				// dp[-1][k][1] = dp[i][0][1] = -infinity

				// dp[i][j][0] = max(dp[-1][j][0], dp[-1][j][1]+prices[i])
				dp[i][j][0] = 0

				// dp[i][j][1] = max(dp[-1][j][1], dp[-1][j][0]-prices[i])
				dp[i][j][1] = -prices[i]

				continue
			}
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
		}
	}

	return dp[n-1][k][0]
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


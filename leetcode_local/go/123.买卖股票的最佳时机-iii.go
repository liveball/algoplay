/*
 * @lc app=leetcode.cn id=123 lang=golang
 *
 * [123] 买卖股票的最佳时机 III
 */

// @lc code=start
func maxProfit(prices []int) int {
	return maxProfit1(prices)
	// return maxProfit2(prices)
}

// 每次 sell 之后要等一天才能继续交易。只要把这个特点融入上一题的状态转移方程即可：

// dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
// dp[i][1] = max(dp[i-1][1], dp[i-2][0] - prices[i])
// 解释：第 i 天选择 `buy` 的时候，要从 i-2 的状态转移，而不是 i-1 。
func maxProfit1(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, 3)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = make([]int, 2)
		}
	}
	dp[0][0][0] = 0
	dp[0][1][1] = -prices[0]
	dp[0][2][0] = 0
	dp[0][2][1] = -prices[0]

	for i := 1; i < n; i++ {
		for k := 2; k >= 1; k-- {
			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
		}
	}

	return dp[n-1][2][0]

}

func maxProfit2(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	k := 2
	dp := make([][][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][]int, k+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = make([]int, 2)
		}
	}

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


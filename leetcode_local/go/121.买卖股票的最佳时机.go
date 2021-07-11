/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */
// @lc code=start
func maxProfit(prices []int) int {
	// return getMax(prices)

	// return getMax2(prices)

	return maxProfitDP(prices)
}

// 解法一 模拟 DP
func maxProfitDP(prices []int) int {
	if len(prices) < 1 {
		return 0
	}
	min, maxProfit := prices[0], 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-min > maxProfit {
			maxProfit = prices[i] - min
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	return maxProfit
}

const (
	MIN = -1 << 31
)

func getMax(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	fmt.Println(MIN)
	dp_i_0 := 0
	dp_i_1 := MIN
	for i := 0; i < n; i++ {
		dp_i_0 = max(dp_i_0, dp_i_1+prices[i])
		dp_i_1 = max(dp_i_1, -prices[i])
	}

	return dp_i_0
}

func getMax2(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	// dp := [n][2]int
	dp := make([][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 2)
	}

	// fmt.Println(dp)

	for i := 0; i < n; i++ {
		if i-1 == -1 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}

	return dp[n-1][0]
}

// dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][1][1] + prices[i])
// dp[i][1][1] = max(dp[i-1][1][1], dp[i-1][0][0] - prices[i])
//             = max(dp[i-1][1][1], -prices[i])
// 解释：k = 0 的 base case，所以 dp[i-1][0][0] = 0。

// 现在发现 k 都是 1，不会改变，即 k 对状态转移已经没有影响了。
// 可以进行进一步化简去掉所有 k：
// dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
// dp[i][1] = max(dp[i-1][1], -prices[i])

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// @lc code=end


/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */
// @lc code=start
func maxProfit(prices []int) int {
	// return maxProfit1(prices)
	return maxProfit2(prices)
}

//dp[i][k][0]=max(dp[i-1][k][0], dp[i-1][k][1]+price[0])
//dp[i][k][1]=max(dp[i-1][k][1], dp[i-1][k][0]-price[0])

// k==1 交易一次
//dp[i][0]=max(dp[i-1][0], dp[i-1][1]+price[0])
//dp[i][1]=max(dp[i-1][1], dp[i-1][0]-price[0])
//        =max(dp[i-1][1], -price[0])
// 解释：k = 0 的 base case，所以 dp[i-1][0][0] = 0。
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
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}

	return dp[n-1][0]
}

func maxProfit2(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	dp := make([]int, 2)

	dp[0] = 0
	dp[1] = -prices[0]

	for i := 1; i < n; i++ {
		dp[0] = max(dp[0], dp[1]+prices[i])
		dp[1] = max(dp[1], -prices[i])
	}

	return dp[0]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// @lc code=end


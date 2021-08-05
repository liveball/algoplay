/*
 * @lc app=leetcode.cn id=188 lang=golang
 *
 * [188] 买卖股票的最佳时机 IV
 */

// @lc code=start
func maxProfit(k int, prices []int) int {
	return maxProfit1(k, prices)
}

// 有了上一题 k = 2 的铺垫，这题应该和上一题的第一个解法没啥区别。
// 但是出现了一个超内存的错误，原来是传入的 k 值会非常大，
// dp 数组太大了。现在想想，交易次数 k 最多有多大呢？

// 一次交易由买入和卖出构成，至少需要两天。所以说有效的限制 k 应该不超过 n/2，
// 如果超过，就没有约束作用了，相当于 k = +infinity。这种情况是之前解决过的。

// 直接把之前的代码重用：
func maxProfit1(maxK int, prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	if maxK > n/2 { // k = +infinity。这种情况是之前解决过的。
		return maxProfit_k_inf(prices)
	}

	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, maxK+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = make([]int, 2)
		}
	}

	for i := 0; i < n; i++ {
		for k := maxK; k >= 1; k-- {
			if i-1 == -1 {
				dp[i][k][0] = 0
				dp[i][k][1] = -prices[i]
				continue
			}

			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
		}
	}

	return dp[n-1][maxK][0]
}

func maxProfit_k_inf(prices []int) int {
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

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// @lc code=end


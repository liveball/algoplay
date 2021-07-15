/*
 * @lc app=leetcode.cn id=518 lang=golang
 *
 * [518] 零钱兑换 II
 */

// @lc code=start
func change(amount int, coins []int) int {
	// return change1(amount, coins)
	return change2(amount, coins) //状态压缩
}

func change1(amount int, coins []int) int {
	n := len(coins)
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, amount+1)
	}

	for i := 0; i <= n; i++ {
		dp[i][0] = 1 //如果凑出的目标金额为0，则表示是唯一的凑法
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= amount; j++ {
			if j-coins[i-1] >= 0 {
				//如果你把这第 i 个物品装入了背包
				// 也就是说你使用 coins[i] 这个面值的硬币，
				// 那么 dp[i][j] 应该等于 dp[i][j-coins[i-1]]。

				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			} else {
				//如果你不把这第 i 个物品装入背包
				// 也就是说你不使用 coins[i] 这个面值的硬币，
				// 那么凑出面额 j 的方法数 dp[i][j] 应该等于 dp[i-1][j]，继承之前的结果。
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n][amount]
}

//状态压缩
func change2(amount int, coins []int) int {
	n := len(coins)
	dp := make([]int, amount+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		for j := 1; j <= amount; j++ {
			if j-coins[i-1] >= 0 {
				dp[j] = dp[j] + dp[j-coins[i-1]]
			}
		}
	}

	return dp[amount]
}

// @lc code=end


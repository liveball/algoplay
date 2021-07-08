/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

package solution

// @lc code=start
func coinChange(coins []int, amount int) int {
	if amount < 1 {
		return 0
	}

	// Time Limit Exceeded 77/182 cases passed (N/A)
	// return recursion(0, coins, amount)

	// return dpUpToDown(coins, amount, make([]int, amount))

	return dpDownToUp(coins, amount)
}

func recursion(idxCoin int, coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	if idxCoin < len(coins) && amount > 0 {
		maxVal := amount / coins[idxCoin]
		minCost := IntMax
		for x := 0; x <= maxVal; x++ {
			if x*coins[idxCoin] <= amount {
				res := recursion(idxCoin+1, coins, amount-x*coins[idxCoin])
				if res != -1 {
					minCost = min(minCost, res+x)
				}
			}
		}

		if minCost == IntMax {
			return -1
		} else {
			return minCost
		}
	}

	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

const (
	IntMax = 1<<31 - 1
)

// Accepted
// 182/182 cases passed (40 ms)
// Your runtime beats 20.06 % of golang submissions
// Your memory usage beats 20.34 % of golang submissions (6.8 MB)

func dpUpToDown(coins []int, n int, count []int) int {
	if n < 0 {
		return -1
	}
	if n == 0 {
		return 0
	}
	if count[n-1] != 0 {
		return count[n-1]
	}

	res := IntMax
	for _, coin := range coins {
		sub := dpUpToDown(coins, n-coin, count)
		if sub >= 0 && sub < res {
			res = sub + 1
		}
	}

	if res == IntMax {
		count[n-1] = -1
	} else {
		count[n-1] = res
	}

	return count[n-1]
}

// Accepted
// 182/182 cases passed (8 ms)
// Your runtime beats 92.62 % of golang submissions
// Your memory usage beats 67.83 % of golang submissions (6.1 MB)
func dpDownToUp(coins []int, amount int) int {
	max := amount + 1
	dp := make([]int, amount+1)
	for k := range dp {
		dp[k] = max
	}
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	} else {
		return dp[amount]
	}
}

// @lc code=end

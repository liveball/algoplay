/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 */

// @lc code=start
func minPathSum(grid [][]int) int {
	// return minPathSum1(grid)

	return minPathSum2(grid)
}

func minPathSum2(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	if m == 0 || n == 0 {
		return 0
	}

	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}

	for j := 1; j < n; j++ {
		grid[0][j] += grid[0][j-1]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] += min(grid[i][j-1], grid[i-1][j])
		}
	}

	return grid[m-1][n-1]

}

func minPathSum1(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	if m == 0 || n == 0 {
		return 0
	}

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < len(dp); i++ {
		if i == 0 {
			dp[i][0] = grid[i][0]
		} else {
			dp[i][0] = dp[i-1][0] + grid[i][0]
		}
	}

	for i := 0; i < len(dp[0]); i++ {
		if i == 0 {
			dp[0][i] = grid[0][i]
		} else {
			dp[0][i] = dp[0][i-1] + grid[0][i]
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i][j-1], dp[i-1][j]) + grid[i][j]
		}
	}

	return dp[m-1][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// @lc code=end


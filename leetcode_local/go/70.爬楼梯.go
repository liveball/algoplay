/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	//纯递归
	// return climbStairs1(n)

	//使用备忘录
	// res := make(map[int]int)
	// return climbStairs2(n, res)

	//使用动态规划
	// return climbStairs3(n)

	//迭代法，空间压缩
	return climbStairs4(n)
}

//递归，超出限制
func climbStairs1(n int) int {
	if n <= 2 {
		return n
	}

	return climbStairs1(n-1) + climbStairs1(n-2)
}

//递归，使用备忘录，防止重复计算
func climbStairs2(n int, res map[int]int) int {
	if n <= 2 {
		return n
	}

	if v, ok := res[n]; ok {
		return v
	}

	res[n] = climbStairs2(n-1, res) + climbStairs2(n-2, res)
	return res[n]
}

//动态规划，迭代法
func climbStairs3(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

//迭代法，空间压缩
func climbStairs4(n int) int {
	if n <= 2 {
		return n
	}

	a := 1
	b := 2
	sum := 0
	for i := 3; i <= n; i++ {
		sum = a + b
		a = b
		b = sum
	}

	return sum
}

// @lc code=end


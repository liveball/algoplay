/*
 * @lc app=leetcode.cn id=509 lang=golang
 *
 * [509] 斐波那契数
 */

// @lc code=start
func fib(n int) int {
	// return fibRecursion(n) //递归

	//带有备忘录的递归
	// var memoMap = make(map[int]int)
	// memoMap[1] = 1
	// memoMap[2] = 1
	// return fibMemo(memoMap, n) //备忘录

	//动态规划
	// return fibDp(n)

	//空间复杂度为o(1)解法
	return fibConstSpace(n)
}

func fibConstSpace(n int) int {
	if n < 2 {
		return n
	}

	prev := 0
	cur := 1
	for i := 0; i < n-1; i++ {
		sum := prev + cur
		prev = cur
		cur = sum
	}

	return cur
}

func fibDp(n int) int {
	dp := make([]int, n)
	dp[1] = 1
	dp[2] = 1
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func fibMemo(memoMap map[int]int, n int) int {
	if n > 0 && memoMap[n] == 0 { //未计算过则缓存
		memoMap[n] = fibMemo(memoMap, n-1) + fibMemo(memoMap, n-2)
	}

	return memoMap[n]

}

//爆栈
func fibRecursion(n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	return fibRecursion(n-1) + fibRecursion(n-2)
}

// Line 5: panic.go, line 1117
// runtime: out of memory: cannot allocate 268435456-byte block (201064448 in use)
// fatal error: out of memory
// runtime stack:
// runtime.throw(0x4f0f52, 0xd)
// panic.go, line 1117
// runtime.stackalloc(0x10000000, 0xc0080e0000, 0x0)
// stack.go, line 401
// runtime.copystack(0xc000000180, 0x10000000)
// stack.go, line 849

// @lc code=end


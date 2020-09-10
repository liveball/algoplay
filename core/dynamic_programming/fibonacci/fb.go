package main

//递归
func fbRecursion(n int) (res int) {
	if n == 1 || n == 2 {
		res = 1
		return res
	}

	res = fbRecursion(n-1) + fbRecursion(n-2)
	return
}

//备忘录
func fbMemo(n int, tmap map[int]int) int {
	if n == 1 || n == 2 {
		return 1
	}

	if v, ok := tmap[n]; ok {
		return v
	} else {
		v = fbMemo(n-1, tmap) + fbMemo(n-2, tmap)
		tmap[n] = v
		return v
	}
}

func fbDp(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 1

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

func fbDp2(n int) int {
	a := 1
	b := 1

	sum := 0
	for i := 3; i <= n; i++ {
		sum = a + b
		a = b
		b = sum
	}

	return sum
}

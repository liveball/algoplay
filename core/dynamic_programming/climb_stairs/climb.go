package main

import (
	"fmt"
	"time"
)

func main() {
	var stairs int
	start := time.Now()
	//stairs = climb(40) //313.01203ms 165580141
	//stairs = climbMemo(40,make(map[int]int))//30.123µs 165580141
	stairs = climbByDynamic(40) //188ns 165580141
	//stairs = climbStairs(40) //2.843µs 165580141

	end := time.Now()
	fmt.Println(end.Sub(start), stairs)
}

func climb(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 { //1个台阶
		return 1
	}

	if n == 2 { //两个台阶
		return 2
	}

	return climb(n-1) + climb(n-2) //n-2个台阶
}

func climbMemo(n int, tmap map[int]int) int {
	if n < 1 {
		return 0
	}
	if n == 1 { //1个台阶
		return 1
	}

	if n == 2 { //两个台阶
		return 2
	}

	if v, ok := tmap[n]; ok {
		return v
	} else {
		v = climbMemo(n-1, tmap) + climbMemo(n-2, tmap)
		tmap[n] = v
		return v
	}
}

//动态规划算法
func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2

	for i := 3; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

func climbByDynamic(n int) int {
	t, a, b := 0, 1, 2

	for i := 3; i <= n; i++ {
		t = a + b
		a = b
		b = t
	}

	return t
}

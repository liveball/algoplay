/*
 * @lc app=leetcode.cn id=441 lang=golang
 *
 * [441] 排列硬币
 */

// @lc code=start
func arrangeCoins(n int) int {
	// return arrangeCoins1(n)//模拟
	return arrangeCoins2(n) //数学公式
}

//模拟
func arrangeCoins1(n int) int {
	k := 1
	for n >= k {
		n = n - k
		k++
		// fmt.Println(111, n, k)
	}

	return k - 1
}

//数学公式 (1+x)*x/2=n x=floor(sqrt(2*n+1/4) - 1/2)
func arrangeCoins2(n int) int {
	if n <= 0 {
		return 0
	}

	x := math.Sqrt(2*float64(n)+0.25) - 0.5
	return int(x)
}

// @lc code=end


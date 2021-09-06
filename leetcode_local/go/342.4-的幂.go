/*
 * @lc app=leetcode.cn id=342 lang=golang
 *
 * [342] 4的幂
 */

// @lc code=start
func isPowerOfFour(n int) bool {
	// return isPowerOfFour1(n)
	return isPowerOfFour2(n)
}

func isPowerOfFour1(n int) bool {
	for n >= 4 {
		if n%4 == 0 {
			n = n / 4
		} else {
			return false
		}
	}

	return n == 1
}

func isPowerOfFour2(n int) bool {
	return n > 0 && (n&(n-1) == 0) && (n-1)%3 == 0
}

// @lc code=end


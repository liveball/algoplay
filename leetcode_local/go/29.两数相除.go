/*
 * @lc app=leetcode.cn id=29 lang=golang
 *
 * [29] 两数相除
 */

// @lc code=start
func divide(dividend int, divisor int) int {
	return divide1(dividend, divisor)
}

func divide1(dividend int, divisor int) int {
	sign, res := -1, 0
	// low, high := 0, abs(dividend)

	if dividend == 0 {
		return 0
	}

	if divisor == 1 {
		return dividend
	}

	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	if dividend > 0 && divisor > 0 || dividend < 0 && divisor < 0 {
		sign = 1
	}

	if dividend > math.MaxInt32 {
		dividend = math.MaxInt32
	}

	res = binarySearchQuotient(
		0,
		abs(dividend),
		abs(divisor),
		abs(dividend),
	)
	if res > math.MaxInt32 {
		return sign * math.MaxInt32
	}

	if res < math.MinInt32 {
		return sign * math.MinInt32
	}

	return sign * res
}

func abs(n int) int {
	y := n >> 31
	return (n ^ y) - y
}

func binarySearchQuotient(low, high, val, dividend int) int {
	quotient := low + (high-low)>>1

	if ((quotient+1)*val > dividend && quotient*val <= dividend) ||
		((quotient+1)*val >= dividend && quotient*val < dividend) {
		if (quotient+1)*val == dividend {
			return quotient + 1
		}

		return quotient
	}

	if (quotient+1)*val > dividend && quotient*val > dividend {
		return binarySearchQuotient(low, quotient-1, val, dividend)
	}

	if (quotient+1)*val < dividend && quotient*val < dividend {
		return binarySearchQuotient(quotient+1, high, val, dividend)
	}

	return 0
}

// @lc code=end


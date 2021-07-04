/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	cur := 0
	for x > cur {
		cur = cur*10 + x%10
		x = x / 10

		fmt.Println(11, cur, x)
	}

	return x == cur || x == cur/10
}

// @lc code=end


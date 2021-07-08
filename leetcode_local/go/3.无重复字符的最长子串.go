/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	return slidingWindow(s)
}

func slidingWindow(s string) int {
	window := make(map[rune]int)
	left, right := 0, 0
	size := len(s)

	res := 0
	for right < size {
		c := rune(s[right])
		right++
		window[c]++

		for window[c] > 1 { //判断左边窗口是否收缩
			d := rune(s[left])
			left++
			window[d]--
		}
		res = max(res, right-left)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// @lc code=end


/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 */

// @lc code=start
func findAnagrams(s string, p string) []int {
	return slidingWindow(p, s)
}

func slidingWindow(s1, s2 string) []int {
	need, window := make(map[rune]int), make(map[rune]int)
	left, right := 0, 0
	size1, size2, valid := len(s1), len(s2), 0
	for _, v := range s1 {
		need[v]++
	}

	res := make([]int, 0)
	for right < size2 {
		c := rune(s2[right])
		right++

		if _, ok := need[c]; ok { //向右滑动，递增窗口
			window[c]++
			if need[c] == window[c] {
				valid++
			}
		}

		for right-left >= size1 { //只要窗口内字符串长度大于s1长度
			if valid == len(need) {
				res = append(res, left)
			}

			d := rune(s2[left])
			left++

			if _, ok := need[d]; ok {
				if need[d] == window[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	return res
}

// @lc code=end


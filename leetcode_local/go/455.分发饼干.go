/*
 * @lc app=leetcode.cn id=455 lang=golang
 *
 * [455] 分发饼干
 */

// @lc code=start
func findContentChildren(g []int, s []int) int {
	if len(g) == 0 || len(s) == 0 {
		return 0
	}

	sort.Ints(g)
	sort.Ints(s)
	res := 0
	index := len(s) - 1 //饼干数组下标
	for i := len(g) - 1; i >= 0; i-- {
		if index >= 0 && s[index] >= g[i] {
			res++
			index--
		}
	}

	return res
}

// @lc code=end


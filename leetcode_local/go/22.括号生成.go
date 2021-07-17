/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) []string {
	res := []string{}
	if n == 0 {
		return res
	}

	dfs(n, n, "", &res)
	return res
}

func dfs(l, r int, str string, res *[]string) {
	if l == 0 && r == 0 {
		*res = append(*res, str)
		return
	}

	if l > 0 {
		dfs(l-1, r, str+"(", res)
	}

	if r > 0 && l < r {
		dfs(l, r-1, str+")", res)
	}

}

// @lc code=end


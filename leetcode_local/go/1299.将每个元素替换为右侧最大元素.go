/*
 * @lc app=leetcode.cn id=1299 lang=golang
 *
 * [1299] 将每个元素替换为右侧最大元素
 */

// @lc code=start
func replaceElements(arr []int) []int {
	tmp, m := 0, -1
	for i := len(arr) - 1; i >= 0; i-- {
		tmp = arr[i]
		arr[i] = m
		m = max(m, tmp)
	}

	return arr
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end


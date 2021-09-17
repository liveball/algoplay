/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	// return strings.Index(haystack, needle)

	if needle == "" {
		return 0
	}
	m := len(haystack)
	n := len(needle)

	for i := range haystack {
		if i+n <= m {
			if haystack[i:i+n] == needle {
				return i
			}
		}

	}

	return -1
}

// @lc code=end


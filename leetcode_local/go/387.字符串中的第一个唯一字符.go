/*
 * @lc app=leetcode.cn id=387 lang=golang
 *
 * [387] 字符串中的第一个唯一字符
 */

// @lc code=start
func firstUniqChar(s string) int {
	return firstUniqChar1(s)
	// return firstUniqChar2(s)
}
func firstUniqChar1(s string) int {
	var arr [26]int
	for i, v := range s {
		arr[v-'a'] = i
	}

	for i, v := range s {
		if i == arr[v-'a'] {
			return i
		} else {
			arr[v-'a'] = -1
		}
	}

	return -1
}

func firstUniqChar2(s string) byte {
	if s == "" {
		return ' '
	}

	var chs [26]int
	for _, v := range s {
		chs[v-'a']++
	}

	for k, v := range s {
		if chs[v-'a'] == 1 {
			return s[k]
		}
	}

	return ' '
}

// @lc code=end


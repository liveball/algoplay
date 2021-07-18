/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 */

// @lc code=start
func isPalindrome(s string) bool {
	// return isPalindrome1(s)
	return isPalindrome2(s)
}
func isPalindrome1(s string) bool {
	// 空是有效的回文
	if s == "" {
		return true
	}
	// 统一转换为小写
	s = strings.ToLower(s)
	// 利用strings builder来构建string
	var build strings.Builder
	// 只要 a-z 和 0-9的字符
	for i := range s {
		if ('a' <= s[i] && s[i] <= 'z') || ('0' <= s[i] && s[i] <= '9') {
			build.WriteByte(s[i])
		}
	}
	// 得到结果
	s = build.String()

	i := 0
	j := len(s) - 1

	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func isPalindrome2(s string) bool {
	s = strings.ToLower(s)
	i, j := 0, len(s)-1
	for i < j {
		for i < j && !isChar(s[i]) {
			i++
		}
		for i < j && !isChar(s[j]) {
			j--
		}

		if s[i] != s[j] {
			return false
		}

		i++
		j--
	}
	return true
}

// 判断 c 是否是字符或者数字
func isChar(c byte) bool {
	if ('a' <= c && c <= 'z') || ('0' <= c && c <= '9') {
		return true
	}
	return false
}

// @lc code=end


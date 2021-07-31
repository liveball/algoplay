/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	return isValid1(s)
}

func isValid1(s string) bool {
	if s == "" {
		return true
	}

	stack := make([]rune, 0, len(s))
	checkMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, v := range s {
		a, ok := checkMap[v]
		if !ok { //如果不是右括号，则是左括号，入栈
			stack = append(stack, v) //注意这里是左括号入栈，循环变量v
		} else if len(stack) > 0 && stack[len(stack)-1] == a { //如果栈不空，并且栈顶元素是右括号，出栈
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}

// @lc code=end


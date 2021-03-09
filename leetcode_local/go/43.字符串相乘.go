/*
 * @lc app=leetcode.cn id=43 lang=golang
 *
 * [43] 字符串相乘
 */

// @lc code=start
func multiply(num1 string, num2 string) string {
	m, n := len(num1), len(num2)
	if m == 0 || n == 0 {
		return "0"
	}

	res := make([]int, m+n)

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			mul := int(num1[i]-'0') * int(num2[j]-'0')
			p1 := i + j
			p2 := i + j + 1
			sum := res[p2] + mul
			res[p2] = sum % 10
			res[p1] += sum / 10
		}
	}

	i := 0
	for i < len(res) && res[i] == 0 {
		i++
	}

	str := ""
	for ; i < len(res); i++ {
		str += strconv.Itoa(res[i])
	}
	if str == "" {
		str = "0"
	}
	return str
}

// @lc code=end


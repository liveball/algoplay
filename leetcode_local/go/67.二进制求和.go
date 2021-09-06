/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 */

// @lc code=start
func addBinary(a string, b string) string {
	if len(b) > len(a) {
		a, b = b, a
	}

	res := make([]string, len(a)+1)
	i, j, k := len(a)-1, len(b)-1, len(a)
	c := 0
	for i >= 0 && j >= 0 {
		ai, _ := strconv.Atoi(string(a[i]))
		bj, _ := strconv.Atoi(string(b[j]))

		sum := ai + bj + c
		res[k] = strconv.Itoa(sum % 2)
		c = sum / 2
		i--
		j--
		k--
	}

	for i >= 0 {
		ai, _ := strconv.Atoi(string(a[i]))
		sum := ai + c
		res[k] = strconv.Itoa(sum % 2)
		c = sum / 2
		i--
		k--
	}

	if c > 0 {
		res[k] = strconv.Itoa(c)
	}

	ret := strings.Join(res, "")
	// fmt.Println(11, ret)
	return ret
}

// @lc code=end


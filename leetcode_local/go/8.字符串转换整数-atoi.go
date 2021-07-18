/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
func myAtoi(s string) int {
	// return strToInt1(s)
	return strToInt2(s)
}

const (
	intMin = -1 << 31
	intMax = 1<<31 - 1
)

func strToInt1(str string) int {
	if str == "" {
		return 0
	}

	i := 0
	for i < len(str)-1 && str[i] == ' ' {
		i++
	}

	sign := 1
	if str[i] == '-' {
		sign = -1
		i++
	} else if str[i] == '+' {
		i++
	}

	res := 0
	for ; i < len(str) && str[i] >= '0' && str[i] <= '9'; i++ {
		res = res*10 + sign*int(str[i]-'0')

		// fmt.Println(str[i]-'0', res)

		if res > intMax {
			res = intMax
		}

		if res < intMin {
			res = intMin
		}

	}

	return res
}

func strToInt2(str string) int {
	if str == "" {
		return 0
	}

	i := 0                                //记录字符移动位置
	for i < len(str)-1 && str[i] == ' ' { //删除空格
		i++
	}

	sign := 1 //正负数标记
	if str[i] == '-' {
		sign = -1
		i++
	} else if str[i] == '+' {
		i++
	}

	res := 0
	for ; i < len(str) && str[i] >= '0' && str[i] <= '9'; i++ {

		res = res*10 + sign*int((str[i]-'0'))

		fmt.Println(str[i]-'0', res)

		if res > intMax {
			res = intMax
		} else if res < intMin {
			res = intMin
		}
	}

	return res
}

// @lc code=end


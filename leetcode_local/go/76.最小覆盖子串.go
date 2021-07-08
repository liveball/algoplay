/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */

// @lc code=start
func minWindow(s string, t string) string {
	return slidingWindow(s, t)
}

func slidingWindow(s string, t string) (res string) {
	if s == "" {
		res = ""
		return
	}

	need := make(map[rune]int)
	window := make(map[rune]int)
	for _, v := range t {
		need[v]++
		// fmt.Println(k, v)
	}

	left, right := 0, 0
	valid, start := 0, 0
	length := MAXInt

	for right < len(s) {
		//将字符c移入窗口
		c := rune(s[right])
		right++

		if _, ok := need[c]; ok {
			window[c]++

			if need[c] == window[c] {
				valid++
			}
		}

		//已经找到所有目标字符，可以缩小窗口
		for valid == len(need) {
			//更新最小串
			if right-left < length {
				start = left
				length = right - left
			}

			//将字符d移出窗口
			d := rune(s[left])
			left++

			if _, ok := need[d]; ok {
				// fmt.Println(22, left, right, start, length, string(d))

				if need[d] == window[d] {
					valid--
				}
				window[d]--
			}
		}

	}

	// fmt.Println(111, left, right, start, length)
	if length == MAXInt {
		res = ""
	} else {
		res = s[start:][0:length] //注意字符串截取
	}
	return
}

const (
	MAXInt = 1 << 31
)

// @lc code=end


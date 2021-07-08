/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	// return groupAnagramsBySort(strs)

	return groupAnagramsByCount(strs)
}

func groupAnagramsBySort(strs []string) [][]string {
	group := make(map[string][]string)
	for _, str := range strs {
		bs := []byte(str)
		sort.Slice(bs, func(i, j int) bool {
			return bs[i] < bs[j]
		})

		sortedStr := string(bs)
		group[sortedStr] = append(group[sortedStr], str)
	}

	res := make([][]string, 0, len(group))
	for _, v := range group {
		res = append(res, v)
	}

	return res
}

func groupAnagramsByCount(strs []string) [][]string {
	countMap := make(map[[26]int][]string)
	for _, str := range strs {
		cnt := [26]int{}
		for _, v := range str {
			cnt[v-'a']++
		}
		// eat [1 0 0 0 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1 0 0 0 0 0 0]
		// tea [1 0 0 0 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1 0 0 0 0 0 0]
		// fmt.Println(str, cnt)
		countMap[cnt] = append(countMap[cnt], str)
	}

	res := make([][]string, 0, len(countMap))
	for _, v := range countMap {
		res = append(res, v)
	}

	return res
}

func groupAnagramsBySliding(strs []string) [][]string {
	res := make([][]string, 0)
	// for i := 0; i < len(strs); i++ {
	for j := 0; j < len(strs)-1; j++ {

		fmt.Println(strs[j], strs[j+1])
		if slidingWindow(strs[j], strs[j+1]) {
			var t []string
			t = make([]string, 0)
			t = append(t, strs[j])
			t = append(t, strs[j+1])
			res = append(res, t)
		} else {
			var t1, t2 []string
			t1 = append(t1, strs[j])
			t2 = append(t2, strs[j+1])
			res = append(res, t1)
			res = append(res, t2)
		}
	}
	// }

	return res
}

func slidingWindow(s1, s2 string) bool {
	need, window := make(map[rune]int), make(map[rune]int)
	left, right := 0, 0
	size1, size2, valid := len(s1), len(s2), 0
	for _, v := range s1 {
		need[v]++
	}

	for right < size2 {
		c := rune(s2[right])
		right++

		if _, ok := need[c]; ok { //向右滑动，递增窗口
			window[c]++
			if need[c] == window[c] {
				valid++
			}
		}

		for right-left >= size1 { //只要窗口内字符串长度大于s1长度
			if valid == len(need) {
				return true
			}

			d := rune(s2[left])
			left++

			if _, ok := need[d]; ok {
				if need[d] == window[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	return false
}

// @lc code=end


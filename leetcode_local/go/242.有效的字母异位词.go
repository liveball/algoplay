/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	return isAnagram0(s, t)

	// return isAnagram1(s, t)
	// return isAnagram2(s, t)
	// return isAnagram3(s, t)
}

//排序
func isAnagram0(s string, t string) bool {
	s1 := []rune(s)
	sort.Stable(sortRunes(s1))

	t1 := []rune(t)
	sort.Stable(sortRunes(t1))

	return string(s1) == string(t1)
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

//hash 计数
func isAnagram1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	a := make(map[byte]int)
	b := make(map[byte]int)

	for _, value := range s {
		a[byte(value)]++
	}

	for _, value := range t {
		b[byte(value)]++
	}

	for k, v := range a {
		vv, ok := b[k]

		if !ok || (ok && v != vv) { //如果s中有字母在t中不存在，或者s和t相同字母的计数不同则返回false
			return false
		}
	}

	return true
}

//可以利用两个长度都为 26 的字符数组来统计每个字符串中小写字母出现的次数，然后再对比是否相等
func isAnagram2(s string, t string) bool {
	if !(len(s) == len(t)) {
		return false

	}

	var (
		bucket1 [26]int
		bucket2 [26]int
	)

	for _, v := range s {
		bucket1[v-'a']++
	}
	for _, v := range t {
		bucket2[v-'a']++
	}

	fmt.Println(bucket1, bucket2)
	return bucket1 == bucket2
}

//可以只利用一个长度为 26 的字符数组，将出现在字符串 s 里的字符个数加 1，而出现在字符串 t 里的字符个数减 1，最后判断每个小写字母的个数是否都为 0
func isAnagram3(s string, t string) bool {
	if !(len(s) == len(t)) {
		return false
	}

	var chs [26]int
	for _, v := range s {
		chs[v-'a']++
	}

	for _, v := range t {
		chs[v-'a']--
	}

	for _, v := range chs {
		if v != 0 {
			return false
		}
	}

	return true
}

// @lc code=end


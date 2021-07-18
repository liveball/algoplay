/*
 * @lc app=leetcode.cn id=567 lang=golang
 *
 * [567] 字符串的排列
 */

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {
	return checkInclusion1(s1, s2)
	// return checkInclusion2(s1, s2)
}
func checkInclusion1(s1 string, s2 string) bool {
	var freq [256]int
	res := false

	if len(s2) == 0 || len(s2) < len(s1) {
		return res
	}

	//统计s1每个字母出现频次
	for i := 0; i < len(s1); i++ {
		freq[s1[i]-'a']++
	}

	left, right, count := 0, 0, len(s1)

	for right < len(s2) {
		//s2中第一个出现的字母，p总数减1
		if freq[s2[right]-'a'] >= 1 {
			count--
		}
		freq[s2[right]-'a']-- //出现频次减1
		right++               //继续右边滑动

		if count == 0 { //左边界移出不符合规范的元素
			res = true
		}

		//判断每个字母是否都被访问过一遍了
		if right-left == len(s1) {
			if freq[s2[left]-'a'] >= 0 {
				count++
			}
			freq[s2[left]-'a']++
			left++
		}
	}

	return res
}

func checkInclusion2(s1, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	cnt1, cnt2 := [26]int{}, [26]int{}
	for i := range s1 {
		cnt1[s1[i]-'a']++
		cnt2[s2[i]-'a']++
	}
	for i := 0; i < len(s2)-len(s1); i++ {
		if cnt1 == cnt2 {
			return true
		}
		cnt2[s2[i]-'a']--
		cnt2[s2[i+len(s1)]-'a']++
	}
	return cnt1 == cnt2
}

// @lc code=end


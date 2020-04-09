package main

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

func Test_palindrome(t *testing.T) { //回文字符串
	println(palindrome("deed"))
	println(palindrome("leetcode"))
}

func Test_lengthOfLongestSubstring(t *testing.T) { //无重复字符的最长子串

	var tests = []string{
		"abcabcbb", //3
		"bbbbb",    //1
		"pwwkew",   //3
	}

	//for _, v := range tests {
	//	fmt.Println(lengthOfLongestSubstring(v))
	//}

	for _, v := range tests {
		fmt.Println(lengthOfLongestSubstring2(v))
	}

	//for _, v := range tests {
	//	fmt.Println(lengthOfLongestSubstring3(v))
	//}
}

func Test_longestCommonPrefix(t *testing.T) {
	tests := [][]string{
		{
			"flower",
			"flow",
			"flight",
		},
		{
			"dog",
			"racecar",
			"car",
		},
	}

	for _, v := range tests {
		fmt.Println(longestCommonPrefix(v))
	}
}

func palindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//
//示例 1:
//
//输入: "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//示例 2:
//
//输入: "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//示例 3:
//
//输入: "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

//这一题和第 438 题，第 3 题，第 76 题，第 567 题类似，用的思想都是"滑动窗口"。
//
//滑动窗口的右边界不断的右移，只要没有重复的字符，就不断的向右扩大窗口边界。
//一旦出现了重复字符，再需要缩小左边界。直到重复的字符移出了左边界。
//接着又可以开始移动滑动窗口的右边界。

//以此类推，不断的刷新记录的窗口大小。
//最终最大的值就是题目中的所求。

func lengthOfLongestSubstring(s string) int {
	window := 0 //窗口大小

	if len(s) == 0 { //字串长度为0返回窗口大小为0
		return window
	}

	if len(s) == 1 { //字符串长度为1返回窗口大小为1
		window = 1
		return window
	}

	var freq [256]int    //定义256 accii 数组
	left, right := 0, -1 //标记左右边界

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for left < len(s) {
		if right+1 < len(s) && freq[s[right+1]-'a'] == 0 { //判断重复字符，如果没有重复出现
			freq[s[right+1]-'a']++ //标记该字符第一次出现
			right++                //扩大右边界

		} else { // freq[s[right+1]] == 1 //right已经到头
			freq[s[left]-'a']-- //移除左边界重复字符
			left++              //缩小左边界
		}

		maxWindow := right - left + 1   //计算当前窗口大小，右边界-左边界+1
		window = max(window, maxWindow) //比较当前窗口和上次窗口大小，取最大窗口
	}

	return window
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//无重复字符的最长子串
func lengthOfLongestSubstring2(s string) int {
	window, start := 0, 0

	for key := 0; key < len(s); key++ {
		isExist := strings.Index(string(s[start:key]), string(s[key])) //判断当前字符是否在窗口内

		if isExist != -1 { //如果在窗口内，重新计算窗口起始位置
			start = start + 1 + isExist

		} else { //如果不在窗口内，计算当前窗口大小
			max := key - start + 1
			if max > window { //比较当前窗口和上次窗口大小，取最大
				window = max
			}
		}
	}

	return window
}

func lengthOfLongestSubstring3(s string) int {
	var window, start int

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	exist := make(map[uint8]int)
	for i := 0; i < len(s); i++ {

		v := s[i]
		if cnt, ok := exist[v]; ok {
			window = max(window, i-start) //更新窗口大小
			start = max(start, cnt+1)     //更新起始位置
		}

		exist[v] = i
	}

	return max(window, len(s)-start)
}

func longestCommonPrefix(strs []string) string {
	n := len(strs)
	// 输入为空直接返回空
	if n == 0 {
		return ""
		// 只输入一个字符串那么它自己就是最长公共后缀
	} else if n == 1 {
		return strs[0]
		// 如果基准字符串是空，那么直接返回空
	} else if len(strs[0]) == 0 {
		return ""
	}

	// 最长公共前缀再长也不可能比最短的输入字符串还要长对吧。输入源越多，基准字符串不是最短字符串的可能性越大。
	// 因此进行初步优化，求输入源中最短的那个字符串长度，然后把基准字符串按该长度重新截取
	minStrLen := math.MaxInt32
	for i := 0; i < n; i++ {
		if minStrLen > len(strs[i]) {
			minStrLen = len(strs[i])
		}
	}

	prefix := strs[0][0:minStrLen]
	// 标志位，如果为true说明在其它所有字符串中都找到了相同的前缀
	var allFound bool
	for {
		allFound = true
		// 循环其它字符串
		for i := 1; i < n; i++ {
			// 查询是否包含基准字符串，因为求的是前缀，所以index值为0
			if strings.Index(strs[i], prefix) != 0 {
				// 不包含的话削掉基准字符串的尾端以便开展下次查询工作
				prefix = prefix[0 : len(prefix)-1]
				// 标志位置为false说明查询失败
				allFound = false
				// 跳出吧，也没必要再查其它字符串了
				break
			}
		}
		// 如果查询成功，跳出大循环，直接返回当前基准字符串就行了。所以该算法最好的情况就是第一次查询就成功。
		// 如果基准字符串都被砍光了，说明没有最长公共后缀，所以该算法的最差的情况就是最后一个字符串和基准字符串完全没有公共前缀，而中间的字符串都能找到公共前缀，这会导致循环minStrLen * (n - 1)次。
		if allFound || len(prefix) == 0 {
			break
		}
	}
	return prefix
}

//字符串的排列
//给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。
//换句话说，第一个字符串的排列之一是第二个字符串的子串。

//输入: s1 = "ab" s2 = "eidbaooo"
//输出: True
//解释: s2 包含 s1 的排列之一 ("ba").
func checkInclusion(s1 string, s2 string) bool {

}
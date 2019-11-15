package main

import (
	"fmt"
	"strings"
	"testing"
)

func Test_palindrome(t *testing.T) { //回文字符串
	println(palindrome("deed"))
	println(palindrome("leetcode"))
}

func Test_lengthOfLongestSubstring(t *testing.T) { //无重复字符的最长子串

	var tests = []string{
		"abcabcbb",
		"bbbbb",
		"pwwkew",
	}

	for _, v := range tests {
		fmt.Println(lengthOfLongestSubstring(v))
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

//无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	window, start := 0, 0

	for key := 0; key < len(s); key++ {
		isExist := strings.Index(string(s[start:key]), string(s[key]))

		if isExist == -1 {
			if (key - start + 1) > window {
				window = key - start + 1
			}
		} else {
			start = start + 1 + isExist
		}
	}

	return window
}

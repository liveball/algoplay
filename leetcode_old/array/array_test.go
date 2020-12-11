package main

import (
	"fmt"
	"testing"
)

//题目是这样的：两个单词如果包含相同的字母，次序不同，则称为字母易位词(anagram)。
// 例如，“silent”和“listen”是字母易位词，而“apple”和“aplee”不是易位词

//LeetCode 第 242 题：给定两个字符串 s 和 t，编写一个函数来判断 t 是否是 s 的字母异位词。
//
//说明：你可以假设字符串只包含小写字母。
//
//示例 1
//输入: s = 'anagram', t = 'nagaram'
//输出: true
//
//示例 2
//输入: s = 'rat', t = 'car'
//输出: false
//
//字母异位词，也就是两个字符串中的相同字符的数量要对应相等。例如，s 等于 “anagram”，t 等于 “nagaram”，s 和 t 就互为字母异位词。
//
//因为它们都包含有三个字符 a，一个字符 g，一个字符 m，一个字符 n，以及一个字符 r。而当 s 为 “rat”，t 为 “car”的时候，s 和 t 不互为字母异位词。

//解题思路
//
//一个重要的前提“假设两个字符串只包含小写字母”，小写字母一共也就 26 个，因此：
//可以利用两个长度都为 26 的字符数组来统计每个字符串中小写字母出现的次数，然后再对比是否相等；
//可以只利用一个长度为 26 的字符数组，将出现在字符串 s 里的字符个数加 1，而出现在字符串 t 里的字符个数减 1，最后判断每个小写字母的个数是否都为 0。
//按上述操作，可得出结论：s 和 t 互为字母异位词

//TODO 实现一个判断字母异位词的函数
//todo

var (
	strs = [][]string{
		{"anagram", "nagaram"},
		{"rat", "car"},
		{"silent", "listen"},
		{"apple", "aplee"},
	}
)

func Test_242(t *testing.T) {
	fmt.Println("get", getOnlyOneChar("aqabcdbdq"))

	//for _, value := range strs {
	//	fmt.Println(value[0], value[1], isAnagram3(value[0], value[1]))
	//}
}

//hash 计数
func isAnagram(s string, t string) bool {
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

func getOnlyOneChar(s string) string {
	if len(s) == 0 {
		return ""
	}

	var chs [26]int
	var chs2 [26]int32
	for _, v := range s {
		chs[v-'a']++
		chs2[v-'a'] = v
	}

	for k, v := range chs {
		if v == 1 {
			return string(chs2[k])
		}
	}

	return ""
}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(sum("2154365543", "4332656442"), 2154365543+4332656442)
}

//# 高精度加法 #
//1.给两个大整数, 用字符串表示, 比如" 2154365543", "4332656442",
//都可能超过1万位, 写一个函数输出他们之和.
//	需要自己实现加法过程, 不能用某些语言自带的高精度加法函数

func sum(a, b string) string {
	if a == "" {
		return b
	}

	if b == "" {
		return a
	}

	m,n:=len(a), len(b)

	var max int
	if m >  n{
		b = strings.Repeat("0", m - n) + b
		max = m

	} else if m <  n {
		a = strings.Repeat("0", n  - m) + a
		max = n
	} else {
		m=n
	}

	var res string

	var carry int
	for i := max - 1; i >= 0; i-- {
		var val int

		ii, _ := strconv.Atoi(string(a[i]))
		jj, _ := strconv.Atoi(string(b[i]))
		val = (ii + jj + carry) % 10

		carry = (ii + jj + carry) / 10

		res = strconv.Itoa(val) + res
	}

	if carry > 0 {
		res = strconv.Itoa(carry) + res
	}

	return res
}

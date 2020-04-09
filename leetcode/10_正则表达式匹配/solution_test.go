package _10_正则表达式匹配

import (
	"fmt"
	"testing"
)

var (
	ps = [][]string{
		{
			"aab",
			"c*a*b",
		},

		{
			"mississippi",
			"mis*is*p*.",
		},
	}
)

func Test_isMatch(t *testing.T) {

	for _, v := range ps {
		ret := isMatch(v[0], v[1])

		fmt.Println(ret)
	}

}

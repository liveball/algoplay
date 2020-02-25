package string

import (
	"fmt"
	"testing"
)

func Test_palindrome(t *testing.T) {
	println(palindrome("deed"))
	println(palindrome("leetcode"))
}

func Test_reverseString(t *testing.T) {
	a := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(a)
	fmt.Println(string(a))

	b:=[]byte("abcdef")
	reverseString(b)
	fmt.Println(string(b))
}

func Test_onlyOneChar(t *testing.T) {
	fmt.Println(getOnlyOneCharPos("aqabcdbdq"))
	fmt.Println(getOnlyOneChar("aqabcdbdq"))
	fmt.Println(getOnlyOneChar2("aqabcdbdq"))
}

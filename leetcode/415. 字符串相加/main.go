package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(addStrings("0", "9"))

	fmt.Println(addStrings("1", "9"))

	fmt.Println(addStrings("9", "99"))

	fmt.Println(addStrings("1234", "123"))
}

func addStrings(num1 string, num2 string) string {
	res := ""
	carry := 0
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || carry != 0; i, j = i-1, j-1 {
		var x, y int
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}
		sum := x + y + carry
		res = strconv.Itoa(sum%10) + res
		carry = sum / 10
	}
	return res
}

// func addStrings(num1 string, num2 string) string {
// 	 maxLen:=0
// 	 nl1,nl2:=len(num1), len(num2)
// 	 if nl1> nl2 {
// 		 maxLen = nl1
// 	 } else {
// 		 maxLen = nl2
// 	 }

// 	 num1Arr:=make([]int, maxLen)
// 	 num2Arr:=make([]int, maxLen)
// 	 for i:=maxLen-1; i>=0; i--{
//         if nl1<=0{
// 		   num1Arr[i] = 0
// 		} else {
// 			nl1--
// 			num1Arr[i] = int(num1[nl1] - '0')
// 		}
// 	 }

// 	 for i:=maxLen-1; i>=0; i--{
//         if nl2<=0{
// 		   num2Arr[i] = 0
// 		} else {
// 			nl2--
// 			num2Arr[i] = int(num2[nl2] - '0')
// 		}
// 	 }

// 	 carry:=0
// 	 str:=""
// 	 for i:=maxLen-1; i>=0; i--{
// 		sum := num1Arr[i] + num2Arr[i] + carry
// 		t :=sum%10
// 		carry = sum/10
// 		str = strconv.Itoa(t) + str
// 	 }

// 	 if carry!=0 {
// 		str=strconv.Itoa(carry)+str
// 	 }
//      return str
// }

// 415. 字符串相加
// 给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。

// 提示：

// num1 和num2 的长度都小于 5100
// num1 和num2 都只包含数字 0-9
// num1 和num2 都不包含任何前导零
// 你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式

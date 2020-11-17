package main

import "fmt"

var (
	a = []int{4, 2, 5, 7}
)

func main() {
	fmt.Println(sortArrayByParityII(a))
}

//双指针
func sortArrayByParityIIByDoublePointer(A []int) []int {
	if len(A) == 0 {
		return []int{}
	}

	for i, j := 0, 1; i < len(A)-1; i = i + 2 {
		if A[i]%2 == 1 {
			for A[j]%2 == 1 {
				j += 2
			}
			A[i], A[j] = A[j], A[i]
		}
	}

	return A
}

func sortArrayByParityII(A []int) []int {
	n := len(A)
	if n == 0 {
		return []int{}
	}

	res := make([]int, n)
	i := 0
	for _, v := range A {
		if v%2 == 0 {
			res[i] = v
			i += 2
		}
	}

	i = 1
	for _, v := range A {
		if v%2 == 1 {
			res[i] = v
			i += 2
		}
	}

	return res
}

// 给定一个非负整数数组 A， A 中一半整数是奇数，一半整数是偶数。

// 对数组进行排序，以便当 A[i] 为奇数时，i 也是奇数；当 A[i] 为偶数时， i 也是偶数。

// 你可以返回任何满足上述条件的数组作为答案。

//

// 示例：

// 输入：[4,2,5,7]
// 输出：[4,5,2,7]
// 解释：[4,7,2,5]，[2,5,4,7]，[2,7,4,5] 也会被接受。
//

// 提示：

// 2 <= A.length <= 20000
// A.length % 2 == 0
// 0 <= A[i] <= 1000

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/sort-array-by-parity-ii
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

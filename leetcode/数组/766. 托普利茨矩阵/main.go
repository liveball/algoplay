package main

import "fmt"

var (
	matrix = [][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}}

	matrix2 = [][]int{{1, 2}, {2, 2}}

	matrix3 = [][]int{{84}}

	matrix4 = [][]int{{1, 2, 3, 4}, {5, 2, 1, 3}}
	// 1, 2, 3, 4
	// 5, 2, 1, 3
)

func main() {
	fmt.Println(isToeplitzMatrix(matrix))
	fmt.Println(isToeplitzMatrix(matrix2))
	fmt.Println(isToeplitzMatrix(matrix3))

	fmt.Println(isToeplitzMatrix(matrix4))
}

func isToeplitzMatrix(matrix [][]int) bool {
	// m == matrix.length
	// n == matrix[i].length
	// 1 <= m, n <= 20
	// 0 <= matrix[i][j] <= 99

	res := true

	for i := 0; i < len(matrix)-1; i++ {

		for j := 0; j < len(matrix[i])-1; j++ {
			if matrix[i][j] != matrix[i+1][j+1] {
				res = false
			}
		}
	}

	if len(matrix) == 1 {
		res = true
	}

	return res
}

// 给你一个 m x n 的矩阵 matrix 。如果这个矩阵是托普利茨矩阵，返回 true ；否则，返回 false 。

// 如果矩阵上每一条由左上到右下的对角线上的元素都相同，那么这个矩阵是 托普利茨矩阵 。

// 输入：matrix = [[1,2,3,4],[5,1,2,3],[9,5,1,2]]
// 输出：true
// 解释：
// 在上述矩阵中, 其对角线为:
// "[9]", "[5, 5]", "[1, 1, 1]", "[2, 2, 2]", "[3, 3]", "[4]"。
// 各条对角线上的所有元素均相同, 因此答案是 True 。

// 输入：matrix = [[1,2],[2,2]]
// 输出：false
// 解释：
// 对角线 "[1, 2]" 上的元素不同。

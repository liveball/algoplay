/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 */

// @lc code=start
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	res := 0
	numMap := make(map[int]int)
	for _, v := range nums {
		if numMap[v] != 0 {
			continue
		}

		left, right := 0, 0
		sum := 0

		if numMap[v-1] > 0 {
			left = numMap[v-1]
		}

		if numMap[v+1] > 0 {
			right = numMap[v+1]
		}

		sum = left + right + 1
		numMap[v] = sum
		res = max(res, sum)
		// fmt.Println(111, v, res, left, right, sum, numMap)

		numMap[v-left] = sum
		numMap[v+right] = sum

		// fmt.Println(222, v, res, left, right, sum, numMap)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// @lc code=end


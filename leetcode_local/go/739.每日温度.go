/*
 * @lc app=leetcode.cn id=739 lang=golang
 *
 * [739] 每日温度
 */

// @lc code=start
func dailyTemperatures(temperatures []int) []int {
	// return temperatures1(temperatures)
	return temperatures2(temperatures)
}

//O(N^2)
func temperatures1(T []int) []int {
	res := make([]int, len(T))
	for i := 0; i < len(T); i++ {
		for j := i + 1; j < len(T); j++ {
			if T[j] > T[i] {
				res[i] = j - i
				break
			}
		}
	}
	return res
}

//O(N)
func temperatures2(temperatures []int) []int {
	if len(temperatures) == 0 {
		return []int{}
	}

	stack := []int{}
	res := make([]int, len(temperatures))
	for k, v := range temperatures { //O(N)

		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < v { //O(1)
			idx := stack[len(stack)-1]
			res[idx] = k - idx
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, k)
	}

	return res
}

// @lc code=end


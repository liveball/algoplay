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

// 暴力法
// 题意：每个元素找到它右边第一个比它大的元素的位置，求它们的距离
// ii 指向当前元素，jj 扫描它的右边，找到比当前元素大的元素，记录 j-ij−i
// ii 考察下一位，重复上述过程。
// 时间复杂度 O(n^2)
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

// 为什么是栈
// 小项 “出” 了，当前项要 “进”，依然要维持排列的单调性
// 所以当前项要从小项 “出” 的地方 “进”
// 只在一端进出——所以是栈

// 保持单调，为了更快比较
// 如果“数据结构”中的项无序，则新来的项无法快速地比较大小，找到大项
// 如果从小到大排好，则很容易剔除小项，遇到大项

// 复杂度分析
// 一次遍历：O(n)；每个元素都出栈入栈各一次：线性时间的复杂度。综合：O(n)
// 空间复杂度：O(n)
func temperatures2(temperatures []int) []int {
	if len(temperatures) == 0 {
		return []int{}
	}

	stack := []int{}
	res := make([]int, len(temperatures))
	for k, v := range temperatures { //O(N)

		// 1、如果当前元素比栈顶大，则让小项逐个出栈，直到当前元素比栈顶小，停止出栈
		// 2、此时的栈顶元素就是当前项右边的第一个比自己大的元素索引，计算距离
		// 3、当前项入栈
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


/*
 * @lc app=leetcode.cn id=41 lang=golang
 *
 * [41] 缺失的第一个正数
 */

// @lc code=start
func firstMissingPositive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	tm := make(map[int]int)

	for _, v := range nums {
		tm[v] = v
	}

	for i := 1; i < len(nums)+1; i++ {
		if _, ok := tm[i]; !ok {
			return i
		}
	}

	return len(nums) + 1
}

// @lc code=end


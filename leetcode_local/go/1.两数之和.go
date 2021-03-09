/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	targetMap := make(map[int]int)
	for k, v := range nums {
		x := target - v
		if k1, ok := targetMap[x]; ok {
			return []int{k1, k}
		}
		targetMap[v] = k
	}

	return []int{}
}

// @lc code=end


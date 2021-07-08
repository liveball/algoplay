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

	vmap := make(map[int]int)

	for k, v := range nums {
		v1 := target - v
		if k1, ok := vmap[v1]; ok {
			return []int{k1, k}
		}

		vmap[v] = k
	}

	return []int{}
}

// @lc code=end


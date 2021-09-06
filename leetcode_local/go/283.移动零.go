/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}

	// 只扫描数组一遍，不断的用 i，j 标记 0 和非 0 的元素，然后相互交换
	j := 0
	for i, v := range nums {
		//v == 0  i++
		if v != 0 {
			if i != j { //把不等于0的往前面移动
				nums[i], nums[j] = nums[j], nums[i]
				j++
				fmt.Println(i, j, nums)
			} else {
				j++
			}
		}
	}
}

// @lc code=end


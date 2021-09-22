/*
 * @lc app=leetcode.cn id=376 lang=golang
 *
 * [376] 摆动序列
 */

// @lc code=start
func wiggleMaxLength(nums []int) int {
	cnt := len(nums)
	if cnt <= 1 {
		return cnt
	}

	curDiff := 0 //当前差值
	preDiff := 0 //前一对差值
	res := 1     //记录峰值个数，序列默认序列最右边有一个峰值
	for i := 0; i < cnt-1; i++ {
		curDiff = nums[i+1] - nums[i]

		//出现峰值
		if curDiff < 0 && preDiff >= 0 || curDiff > 0 && preDiff <= 0 {
			res++
			preDiff = curDiff
		}
	}

	return res
}

// @lc code=end


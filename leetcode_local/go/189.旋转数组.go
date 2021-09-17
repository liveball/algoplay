/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 */

// @lc code=start
func rotate(nums []int, k int) {
	// rotate1(nums, k)
	rotate2(nums, k)
}

func rotate1(nums []int, k int) {
	n := len(nums)
	newNums := make([]int, n)

	for i, v := range nums {
		newNums[(i+k)%n] = v
	}

	copy(nums, newNums)
}

func rotate2(nums []int, k int) {
	n := len(nums)
	k %= n
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, s, e int) {
	for s < e {
		t := nums[s]
		nums[s] = nums[e]
		nums[e] = t

		s++
		e--
	}
}

// @lc code=end


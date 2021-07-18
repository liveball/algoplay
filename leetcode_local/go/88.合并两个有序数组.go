/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		copy(nums1, nums2)
		return
	}

	i := m - 1
	j := n - 1
	k := m + n - 1

	for k >= 0 {
		if j < 0 || (i >= 0 && nums1[i] > nums2[j]) { //复制num1
			nums1[k] = nums1[i]
			i--
		} else { //复制num2
			nums1[k] = nums2[j]
			j--
		}

		k--
	}
}

// @lc code=end


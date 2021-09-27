/*
 * @lc app=leetcode.cn id=496 lang=golang
 *
 * [496] 下一个更大元素 I
 */

// @lc code=start
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	cnt := len(nums1)
	res := make([]int, 0, cnt)
	if cnt == 0 {
		return res
	}

	checkMap := make(map[int]int)
	for k, v := range nums1 {
		checkMap[v] = k
		res = append(res, -1)
	}

	stack := make([]int, cnt)
	for k, v := range nums2 {
		if v < nums2[stack[len(stack)-1]] {
			stack = append(stack, k)
		} else if v == nums2[stack[len(stack)-1]] {
			stack = append(stack, k)
		} else {
			for len(stack) > 0 && v > nums2[stack[len(stack)-1]] {
				if index, ok := checkMap[nums2[stack[len(stack)-1]]]; ok {
					res[index] = v
				}
				stack = stack[:len(stack)-1]
			}

			stack = append(stack, k)
		}
	}

	return res
}

// @lc code=end


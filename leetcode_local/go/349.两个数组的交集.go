/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	return intersection1(nums1, nums2)
	// return intersect2(nums1, nums2)
}

func intersection1(nums1 []int, nums2 []int) []int {
	m := map[int]bool{}
	var res []int
	for _, n := range nums1 {
		m[n] = true
	}
	for _, n := range nums2 {
		if m[n] {
			delete(m, n)
			res = append(res, n)
		}
	}
	return res
}

func intersection2(nums1, nums2 []int) []int {
	var i, j, k int
	sort.Ints(nums1)
	sort.Ints(nums2)

	cnt := 0
	if len(nums1) > len(nums2) {
		cnt = len(nums1)
	} else {
		cnt = len(nums2)
	}

	nums3 := make([]int, cnt)
	for i < len(nums1) && j < len(nums2) { //反复以下步骤，直到任意一个数组终止
		if nums1[i] < nums2[j] { //如果两个指针的元素不相等，我们将小的一个指针前移。
			i++
		} else if nums1[i] > nums2[j] { //如果两个指针的元素不相等，我们将小的一个指针前移。
			j++
		} else { //设定两个为0的指针，比较两个指针的元素是否相等。如果指针的元素相等，我们将两个指针一起向前移动，并且将相等的元素放入空白数组。
			nums3[k] = nums1[i]
			i++
			j++
			k++
		}
	}

	return nums3[:k]
}

// @lc code=end


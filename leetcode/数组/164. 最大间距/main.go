package main

func main() {

}

// 方法一：基数排序
// 思路与算法

// 一种最简单的思路是将数组排序后再找出最大间距，但传统的基于比较的排序算法（快速排序、归并排序等）
// 均需要 O(N\log N)O(NlogN) 的时间复杂度。我们必须使用其他的排序算法。
// 例如，基数排序可以在 O(N)O(N) 的时间内完成整数之间的排序。

func maximumGap(nums []int) (ans int) {
	n := len(nums)
	if n < 2 {
		return
	}

	buf := make([]int, n)
	maxVal := max(nums...)
	for exp := 1; exp <= maxVal; exp *= 10 {
		cnt := [10]int{}
		for _, v := range nums {
			digit := v / exp % 10
			cnt[digit]++
		}
		for i := 1; i < 10; i++ {
			cnt[i] += cnt[i-1]
		}
		for i := n - 1; i >= 0; i-- {
			digit := nums[i] / exp % 10
			buf[cnt[digit]-1] = nums[i]
			cnt[digit]--
		}
		copy(nums, buf)
	}

	for i := 1; i < n; i++ {
		ans = max(ans, nums[i]-nums[i-1])
	}
	return
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v > res {
			res = v
		}
	}
	return res
}

// 给定一个无序的数组，找出数组在排序之后，相邻元素之间最大的差值。

// 如果数组元素个数小于 2，则返回 0。

// 示例 1:

// 输入: [3,6,9,1]
// 输出: 3
// 解释: 排序后的数组是 [1,3,6,9], 其中相邻元素 (3,6) 和 (6,9) 之间都存在最大差值 3。
// 示例 2:

// 输入: [10]
// 输出: 0
// 解释: 数组元素个数小于 2，因此返回 0。
// 说明:

// 你可以假设数组中所有元素都是非负整数，且数值在 32 位有符号整数范围内。
// 请尝试在线性时间复杂度和空间复杂度的条件下解决此问题。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/maximum-gap
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

package queue

import (
	"fmt"
	"testing"
)

var (
	nums = []int{1, 3, -1, -3, 5, 3, 6, 7}
)

func Test_239(t *testing.T) {
	fmt.Println(maxSlidingWindow(nums, 3))
}

//例题分析
//LeetCode 第 239 题：给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口 k 内的数字，
//滑动窗口每次只向右移动一位。返回滑动窗口最大值。

//示例：给定一个数组以及一个窗口的长度 k，现在移动这个窗口，要求打印出一个数组，数组里的每个元素是当前窗口当中最大的那个数。

//输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
//输出: [3, 3, 5, 5, 6, 7]
//解释:
//
//滑动窗口的位置                最大值
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
//1 [3  -1  -3] 5  3  6  7       3
//1  3 [-1  -3  5] 3  6  7       5
//1  3  -1 [-3  5  3] 6  7       5
//1  3  -1  -3 [5  3  6] 7       6
//1  3  -1  -3  5 [3  6  7]      7

//提示：
//你可以假设 k 总是有效的，在输入数组不为空的情况下，1 ≤ k ≤ 输入数组的大小。

//解题思路
//思路 1：移动窗口，扫描，获得最大值。假设数组里有 n 个元素，算法复杂度就是 O(n)。这是最直观的做法。
//
//思路 2：利用一个双端队列来保存当前窗口中最大那个数在数组里的下标，双端队列新的头就是当前窗口中最大的那个数。
//通过该下标，可以很快地知道新的窗口是否仍包含原来那个最大的数。如果不再包含，我们就把旧的数从双端队列的头删除。
//
//因为双端队列能让上面的这两种操作都能在 O(1) 的时间里完成，所以整个算法的复杂度能控制在 O(n)。

func maxSlidingWindow2(nums []int, k int) []int {
	if len(nums) == 0 || len(nums) < k {
		return make([]int, 0)
	}

	head := 0
	tail := 0
	count := 0

	numsSize := len(nums)
	indexes := make([]int, numsSize)
	results := make([]int, numsSize-k+1)

	for i := 0; i < numsSize; i++ {

		for tail > head && nums[i] >= nums[indexes[tail-1]] {
			tail--
		}

		tail++
		indexes[tail] = i

		if indexes[head] <= i-k {
			head++
		}

		if i >= k-1 {
			count++
			results[count] = nums[indexes[head]]
		}
	}

	return results
}

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || len(nums) < k {
		return make([]int, 0)
	}

	window := make([]int, 0, k) // store the index of nums
	result := make([]int, 0, len(nums)-k+1)

	for i, v := range nums { // if the left-most index is out of window, remove it
		if i >= k && window[0] <= i-k {
			window = window[1:]
		}

		for len(window) > 0 && nums[window[len(window)-1]] < v { // maintain window
			window = window[0 : len(window)-1]
		}

		window = append(window, i) // store the index of nums

		if i >= k-1 {
			result = append(result, nums[window[0]]) // the left-most is the index of max value in nums
		}
	}

	return result
}

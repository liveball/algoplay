/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
	// return maxSlidingWindow1(nums, k)

	//双端队列  O(N)
	return maxSlidingWindow2(nums, k)
	//TODO 优先队列-大顶堆 O(N*logk)
}

func maxSlidingWindow2(nums []int, k int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}

	window := make([]int, 0, k)
	res := make([]int, 0, n-k+1)

	for i, v := range nums {
		if i >= k && window[0] <= i-k { //有元素超越窗口左边界，则移出
			window = window[1:len(window)]
		}

		//如果窗口不空，依次循环窗口里面元素，如果比当前元素小则滑出窗口
		for len(window) > 0 && nums[window[len(window)-1]] < v { //O(1)
			window = window[:len(window)-1]
		}

		//加入窗口
		window = append(window, i)

		//如果当前索引大于窗口左边界，则取窗口最左边的元素
		if i >= k-1 {
			res = append(res, nums[window[0]])
		}
	}

	return res
}

func maxSlidingWindow1(nums []int, k int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}

	res := make([]int, 0, n-k+1)
	for i := 0; i < n-k+1; i++ {
		max := nums[i]
		for j := 1; j < k; j++ {
			if max < nums[i+j] {
				max = nums[i+j]
			}
		}

		res = append(res, max)
	}
	return res
}

// @lc code=end


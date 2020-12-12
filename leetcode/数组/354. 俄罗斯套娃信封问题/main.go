package main

import "sort"

func maxEnvelopes(nums [][]int) int {
	if len(nums) == 0 || len(nums[0]) == 0 {
		return 0
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i][0] < nums[j][0]
	})

	m := len(nums)
	dp := make([]int, m)

	// dp[i]= dp[j]+1 j<i && nums[i]>num[j]
	res := 1
	for i := 0; i < m; i++ {
		dp[i] = 1
		for j := 0; i > 0 && j < i; j++ {
			if nums[i][0] > nums[j][0] && nums[i][1] > nums[j][1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}


// 给定一些标记了宽度和高度的信封，宽度和高度以整数对形式 (w, h) 出现。当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。

// 请计算最多能有多少个信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。

// 说明:
// 不允许旋转信封。

// 示例:

// 输入: envelopes = [[5,4],[6,4],[6,7],[2,3]]
// 输出: 3 
// 解释: 最多信封的个数为 3, 组合为: [2,3] => [5,4] => [6,7]。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/russian-doll-envelopes
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
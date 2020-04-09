package array

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

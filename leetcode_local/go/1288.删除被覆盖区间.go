/*
 * @lc app=leetcode.cn id=1288 lang=golang
 *
 * [1288] 删除被覆盖区间
 */

// @lc code=start
func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}

		return intervals[i][0] < intervals[j][0]
	})

	left := intervals[0][0]
	right := intervals[0][1]
	res := 0
	for i := 1; i < len(intervals); i++ {
		intv := intervals[i]

		//找到覆盖区间
		if left <= intv[0] && right >= intv[1] {
			res++
		}

		//找到相交区间，合并
		if right >= intv[0] && right <= intv[1] {
			right = intv[1]
		}

		//完全不相交， 更新起点和终点
		if right < intv[0] {
			left = intv[0]
			right = intv[1]
		}
	}

	return len(intervals) - res
}

// @lc code=end


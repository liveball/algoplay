/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	// fmt.Println(111, intervals)

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}

		return intervals[i][0] < intervals[j][0]
	})

	// fmt.Println(222, intervals)

	right := intervals[0]
	res := [][]int{}

	i := 1
	for ; i < len(intervals); i++ {
		intv := intervals[i]

		//找到覆盖区间
		if right[1] < intv[0] {
			res = append(res, right)
			// fmt.Println(333, right)
			right = intv
			continue
		}

		//right 吞掉下面一行
		if right[1] > intv[1] {
			continue
		}

		//完全不相交， 更新终点
		if right[1] < intv[1] {
			right[1] = intv[1]
			continue
		}

	}

	if i == len(intervals) { //只有一行
		res = append(res, right)
	}

	return res

}

// @lc code=end


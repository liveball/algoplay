/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	l := len(nums)
	res := make([][]int, 0)
	if l == 0 {
		return res
	}

	path := []int{} // path跟visited是全过程共享的，一定要注意恢复状态
	visited := make(map[int]bool)

	sort.Ints(nums) // 这里是去重的关键逻辑

	dfs(nums, l, 0, path, visited, &res)
	return res
}

func dfs(nums []int, l int, depth int, path []int, visited map[int]bool, res *[][]int) {
	if depth == l {
		*res = append(*res, append([]int{}, path...)) // 保存path当前快照，因为path是共享的，后续会被其他决策修改
		return
	}

	for i, num := range nums {
		if !visited[i] {
			if i > 0 && nums[i] == nums[i-1] && !visited[i-1] { // 这里是去重的关键逻辑
				continue
			}

			visited[i] = true // 更改当前状态，递归调用下一层
			path = append(path, num)
			dfs(nums, l, depth+1, path, visited, res)
			path = path[:len(path)-1] //撤销当前已使用节点
			visited[i] = false        // 每次返回后恢复状态，供下一回合使用
		}
	}
}

// @lc code=end


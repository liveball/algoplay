/*
 * @lc app=leetcode.cn id=337 lang=golang
 *
 * [337] 打家劫舍 III
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
	a, b := dfs(root)
	return max(a, b)
}

func dfs(root *TreeNode) (a, b int) {
	if root == nil {
		return
	}

	l0, l1 := dfs(root.Left)
	r0, r1 := dfs(root.Right)

	//当前节点没有被打劫
	a = max(l0, l1) + max(r0, r1)
	//当前节点被打劫
	b = root.Val + l0 + r0
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// @lc code=end


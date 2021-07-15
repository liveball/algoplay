/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
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
func levelOrder(root *TreeNode) [][]int {
	levels := [][]int{}

	dfs(root, -1, &levels)

	return levels

}

func dfs(root *TreeNode, level int, res *[][]int) {
	if root == nil {
		return
	}

	curLevel := level + 1
	if len(*res) <= curLevel {
		*res = append(*res, []int{})
	}
	(*res)[curLevel] = append((*res)[curLevel], root.Val)
	dfs(root.Left, curLevel, res)
	dfs(root.Right, curLevel, res)
}

// @lc code=end


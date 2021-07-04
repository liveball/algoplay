/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
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
func preorderTraversal(root *TreeNode) []int {
	//根左右
	res := make([]int, 0)
	helper(&res, root)

	return res
}

func helper(res *[]int, root *TreeNode) {
	if root != nil {
		*res = append(*res, root.Val)
		// fmt.Println(root.Val)
		helper(res, root.Left)
		helper(res, root.Right)
	}
}

// @lc code=end


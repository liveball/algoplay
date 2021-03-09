/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
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
func inorderTraversal(root *TreeNode) []int {
	//左根右
	res := make([]int, 0)
	helper(&res, root)

	return res
}

func helper(res *[]int, root *TreeNode) {
	if root != nil {
		helper(res, root.Left)
		*res = append(*res, root.Val)
		// fmt.Println(root.Val)
		helper(res, root.Right)
	}
}

// @lc code=end


/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
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
func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	//分别拉平左右子树
	flatten(root.Left)
	flatten(root.Right)

	//将左子树链接到原来的右子树
	left := root.Left
	right := root.Right
	root.Right = left
	root.Left = nil

	//将原来的右子树链接到右子树
	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = right
}

// @lc code=end


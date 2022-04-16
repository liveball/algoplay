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
	return iterator(root)
	// return recursion(root)
}

func iterator(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node != nil {
			res = append(res, node.Val)
			stack = append(stack, node.Right)
			stack = append(stack, node.Left)
		}
	}

	return res
}

func recursion(root *TreeNode) []int {
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


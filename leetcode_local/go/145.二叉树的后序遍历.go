/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
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
func postorderTraversal(root *TreeNode) []int {
	// return recursion(root)
	return iterator(root)
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
		} else {
			continue
		}

		//前序遍历根左右===>
		// 翻转迭代的左右子树代码根右左 ===>
		//  翻转数组左右根
		stack = append(stack, node.Left)
		stack = append(stack, node.Right)
	}

	return reverse(res)
}

func reverse(arr []int) []int {
	res := make([]int, 0, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		res = append(res, arr[i])
	}
	return res
}

func recursion(root *TreeNode) []int {
	//左右根
	res := make([]int, 0)
	helper(&res, root)

	return res
}

func helper(res *[]int, root *TreeNode) {
	if root != nil {
		helper(res, root.Left)
		helper(res, root.Right)
		*res = append(*res, root.Val)
	}
}

// @lc code=end


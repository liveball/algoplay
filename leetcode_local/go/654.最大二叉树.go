/*
 * @lc app=leetcode.cn id=654 lang=golang
 *
 * [654] 最大二叉树
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
func constructMaximumBinaryTree(nums []int) *TreeNode {
	length := len(nums)
	if length == 0 {
		return nil
	}

	//找出数组中的最大值
	maxVal := -1
	index := 0
	for i := 0; i < length; i++ {
		if nums[i] > maxVal {
			index = i
			maxVal = nums[i]
		}
	}

	node := &TreeNode{
		Val: maxVal,
	}

	//递归调用构造左右子树
	node.Left = constructMaximumBinaryTree(nums[:index])
	node.Right = constructMaximumBinaryTree(nums[index+1:])
	return node
}

// @lc code=end


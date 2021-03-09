/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	inMap := make(map[int]int)
	for k, v := range inorder {
		inMap[v] = k
	}

	return helper(inorder, postorder, inMap, 0, len(inorder)-1, 0, len(postorder)-1)
}

func helper(inorder, postorder []int, inMap map[int]int, inleft, inright, pleft, pright int) *TreeNode {
	if inleft > inright || pleft > pright {
		return nil
	}

	rootKey, ok := inMap[postorder[pright]]
	if !ok {
		return nil
	}

	leftSubNum := rootKey - inleft
	root := &TreeNode{
		Val: postorder[pright],
	}

	root.Left = helper(inorder, postorder, inMap,
		inleft, rootKey-1, pleft, pleft+leftSubNum-1)
	root.Right = helper(inorder, postorder, inMap,
		rootKey+1, inright, pleft+leftSubNum, pright-1)
	return root
}

// @lc code=end


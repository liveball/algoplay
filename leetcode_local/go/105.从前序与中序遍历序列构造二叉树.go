/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	inMap := make(map[int]int)
	for k, v := range inorder {
		inMap[v] = k
	}

	return helper(preorder, inorder, inMap, 0, len(preorder)-1, 0, len(inorder)-1)

}

func helper(preorder, inorder []int, inMap map[int]int, pleft, pright, inleft, inright int) *TreeNode {
	if pleft > pright || inleft > inright {
		return nil
	}

	// 在中序遍历中定位根节点
	// 前序遍历的第一个节点是根节点
	rootKey, ok := inMap[preorder[pleft]]
	if !ok {
		return nil
	}

	root := &TreeNode{Val: preorder[pleft]}
	//计算左子树的数量
	leftSubNum := rootKey - inleft

	//递归构造左子树，并连接到根节点

	// 先序遍历中「从 左边界+1 开始到  左边界+左子树节点数目 」个元素
	// 就对应了中序遍历中「从 左边界 开始到 根节点定位-1」的元素

	root.Left = helper(preorder, inorder, inMap,
		pleft+1, pleft+leftSubNum, inleft, rootKey-1)

	//递归构造右子树，并连接到根节点

	// 先序遍历中「从 左边界+1+左子树节点数目 开始到 右边界」的元素
	// 就对应了中序遍历中「从 根节点定位+1 到 右边界」的元素

	root.Right = helper(preorder, inorder, inMap,
		pleft+1+leftSubNum, pright, rootKey+1, inright)

	return root
}

// @lc code=end


package main

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	Key  int
	Val  int
	Next *Node
	Prev *Node
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

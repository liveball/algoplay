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
	// return recursion(root)
	return iterator(root)
	// return iterator2(root)
}

func iterator(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	stack := make([]*TreeNode, 0)
	res := make([]int, 0)
	cur := root

	for cur != nil || len(stack) > 0 {
		if cur != nil { //一直访问左子树，直到为nil
			stack = append(stack, cur)
			cur = cur.Left

		} else { //弹出最后一个压入栈的左孩子，左右子节点均为nil
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, cur.Val)
			cur = cur.Right //如果为nil 继续从栈中弹出左父节点
		}
	}
	return res
}

func iterator2(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}

	st := list.New()
	cur := root

	for cur != nil || st.Len() > 0 {
		if cur != nil {
			st.PushBack(cur)
			cur = cur.Left
		} else {
			cur = st.Remove(st.Back()).(*TreeNode)
			ans = append(ans, cur.Val)
			cur = cur.Right
		}
	}

	return ans
}

func recursion(root *TreeNode) []int {
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


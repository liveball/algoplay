package binarytree

//前序遍历的递推公式：
//preOrder(r) = print(r) -> preOrder(r->left) ->preOrder(r->right)
//
//中序遍历的递推公式：
//inOrder(r) = inOrder(r->left)-> print(r) -> inOrder(r->right)
//
//后序遍历的递推公式：
//postOrder(r) = postOrder(r->left) -> postOrder(r->right) -> print(r)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	var res []int

	res = append(res, root.Val)
	res = append(res, preOrder(root.Left)...)
	res = append(res, preOrder(root.Right)...)
	return res
}

func preOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	var stack []*TreeNode
	var res []int

	stack = append(stack, root)

	for len(stack) != 0 {
		e := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, e.Val)

		if e.Right != nil {
			stack = append(stack, e.Right)
		}

		if e.Left != nil {
			stack = append(stack, e.Left)
		}
	}
	return res
}

func inOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	res := inOrder(root.Left)
	res = append(res, root.Val)
	res = append(res, inOrder(root.Right)...)
	return res
}

func postOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var res []int
	if root.Left != nil {
		lres := postOrder(root.Left)
		if len(lres) > 0 {
			res = append(res, lres...)
		}
	}

	if root.Right != nil {
		rres := postOrder(root.Right)
		if len(rres) > 0 {
			res = append(res, rres...)
		}
	}

	res = append(res, root.Val)
	return res
}

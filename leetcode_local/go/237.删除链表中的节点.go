/*
 * @lc app=leetcode.cn id=237 lang=golang
 *
 * [237] 删除链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
	if node == nil {
		return
	}

	// 1、其实就是把后面的结点都覆盖上来即可。
	cur := node
	i := 0
	for cur.Next.Next != nil {
		fmt.Println(111, cur)
		cur.Val = cur.Next.Val //覆盖为下个节点的值
		cur = cur.Next         //当前节点指向下一个节点
		i++
		fmt.Println(222, i, cur)
	}
	cur.Val = cur.Next.Val //下下个节点的值覆盖下个节点的值
	cur.Next = nil         //下下下个节点指向nil

	// 2、或者直接当前结点的值等于下一个结点，Next指针指向下下个结点，这样做也可以，
	// 只不过中间有一个结点不被释放，内存消耗多一些。
	// cur := node
	// cur.Val = cur.Next.Val
	// cur.Next = cur.Next.Next
}

// @lc code=end


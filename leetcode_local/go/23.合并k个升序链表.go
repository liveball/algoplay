/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个升序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)

	if length == 0 {
		return nil
	}

	if length == 1 {
		return lists[0]
	}

	mid := length / 2
	left := mergeKLists(lists[:mid])
	right := mergeKLists(lists[mid:])

	return mergeKLists1(left, right)
}

func mergeKLists1(l1, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	dummy := &ListNode{}
	cur := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}

		cur = cur.Next
	}

	if l1 != nil {
		cur.Next = l1
	}

	if l2 != nil {
		cur.Next = l2
	}

	return dummy.Next

}

// @lc code=end


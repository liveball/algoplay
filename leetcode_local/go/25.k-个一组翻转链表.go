/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var dummy ListNode
	pre := &dummy
	dummy.Next = head

	count := 0
	for p := head; p != nil; p = p.Next {
		count++
		if count%k != 0 {
			continue
		}

		cur := pre.Next

		for i := 1; i < k; i++ {
			next := cur.Next
			cur.Next = next.Next
			next.Next = pre.Next
			pre.Next = next
		}
		p = cur
		pre = p
	}

	return dummy.Next
}

// @lc code=end


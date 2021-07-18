/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	// 使用快慢指针寻找链表中点，然后反转后半部分，
	// 再分别从开头和反转后的节点处遍历比较

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var pre, cur *ListNode = nil, slow
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	thead, mid := head, pre
	for thead != nil && mid != nil {
		if thead.Val != mid.Val {
			return false
		}

		thead = thead.Next
		mid = mid.Next
	}

	return true
}

// @lc code=end


/*
 * @lc app=leetcode.cn id=328 lang=golang
 *
 * [328] 奇偶链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	//奇数
	oddHead := &ListNode{}
	odd := oddHead

	evenHead := &ListNode{}
	even := evenHead

	count := 1
	for head != nil {
		if count%2 == 1 { //奇数
			odd.Next = head
			odd = odd.Next
		} else {
			even.Next = head
			even = even.Next
		}

		count++
		head = head.Next
	}

	even.Next = nil //偶数先用光？
	odd.Next = evenHead.Next
	return oddHead.Next
}

// @lc code=end


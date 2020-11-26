package main

func main()  {
	
}

/**
 * Definition for singly-linked list.
 */

 type ListNode struct {
	Val int
   Next *ListNode
}

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

		begin := pre.Next

		for i := 1; i < k; i++ {
			next := begin.Next
			begin.Next = next.Next
			next.Next = pre.Next
			pre.Next = next
		}
		p = begin
		pre = p
	}

	return dummy.Next
}
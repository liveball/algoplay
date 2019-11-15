package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 判断链表是否有环，不能使用额外的空间。
// 给 2 个指针，一个指针是另外一个指针的下一个指针。
// 快指针一次走 2 格，慢指针一次走 1 格。
// 如果存在环，那么前一个指针一定会经过若干圈之后追上慢的指针。

func hasCycle(head *ListNode) bool {
	fast := head
	slow := head
	for slow != nil && fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

func reverseList(head *ListNode) *ListNode {
	var p, q *ListNode

	p = head
	q = head.Next

	head.Next = nil

	for q != nil {
		r := q.Next
		q.Next = p
		p = q
		q = r
	}

	return p
}

func swapList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var dummy ListNode
	dummy.Next = head

	pre := &dummy
	p := dummy.Next
	next := p.Next

	for p != nil && next != nil {
		pre.Next = next
		p.Next = next.Next
		next.Next = p
		pre = p
		p = p.Next
		if p != nil {
			next = p.Next
		}

	}

	return dummy.Next
}

// convert *ListNode to []int
func l2s(head *ListNode) []int {
	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

// convert []int to *ListNode
func s2l(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	res := &ListNode{
		Val: nums[0],
	}
	temp := res
	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode{
			Val: nums[i],
		}
		temp = temp.Next
	}
	return res
}

package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(nil))

	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next = &ListNode{Val: 2}
	head.Next = &ListNode{Val: 1}
	fmt.Println(isPalindrome(head))
}

type ListNode struct {
	Next *ListNode
	Val  int
}

// 使用快慢指针寻找链表中点，然后反转后半部分，再分别从开头和反转后的节点处遍历比较
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

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

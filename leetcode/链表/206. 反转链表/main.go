package main

func main() {

}

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

//iteration
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head
	var pre *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}

//Recursion
// func reverseList(head *ListNode) *ListNode {
// 	if head==nil ||  head.Next==nil{
// 		return head
// 	}

// 	cur:=reverseList(head.Next)
// 	head.Next.Next=head
// 	head.Next=nil
// 	return cur
// }

// 206. 反转链表
// 反转一个单链表。

// 示例:

// 输入: 1->2->3->4->5->NULL
// 输出: 5->4->3->2->1->NULL
// 进阶:
// 你可以迭代或递归地反转链表。你能否用两种方法解决这道题？

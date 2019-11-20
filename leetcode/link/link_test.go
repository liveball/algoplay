package link

import (
	"fmt"
	"testing"
)

const (
	mutexLocked      = 1 << iota // mutex is locked  //1
	mutexWoken                   //2
	mutexStarving                //4
	mutexWaiterShift = iota      //3
)

func Test_24(t *testing.T) {
	fmt.Println(mutexLocked, mutexWoken, mutexStarving, mutexWaiterShift)

	nodes := getNode()
	//printNode(nodes)
	printNode(swapPairs(nodes))

	//printNode(swapPairs2(nodes))
}

func Test_25(t *testing.T) {
	nodes := getNode()
	//printNode(nodes)

	//printNode(reverseKGroup(nodes, 5))
	//当 k=2 时，应当返回：2-›1-›4-›3-›5
	//当 k=3 时，应当返回：3-›2-›1-›4-›5

	printNode(reverse(nodes))

}

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
//你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换

//示例:
//给定 1->2->3->4, 你应该返回 2->1->4->3.
func swapPairs(head *ListNode) *ListNode { //递归交换

	if head == nil || head.Next == nil {
		return head
	}

	t := head.Next                //第2个节点作为新的头节点
	head.Next = swapPairs(t.Next) //递归第二个节点之后的所有节点
	t.Next = head                 //第2个节点指向头节点

	return t
}

func swapPairs2(head *ListNode) *ListNode { //递归交换

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

//LeetCode 第 25 题：给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
//k 是一个正整数，它的值小于或等于链表的长度。
//如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

//示例：
//给定这个链表：1-›2-›3-›4-›5
//当 k=2 时，应当返回：2-›1-›4-›3-›5
//当 k=3 时，应当返回：3-›2-›1-›4-›5
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

func getNode() *ListNode {
	listNode := new(ListNode)
	listNode.Val = 1

	t := listNode
	for i := listNode.Val + 1; i <= 5; i++ {
		t.Next = &ListNode{
			Val: i,
		}

		t = t.Next
	}

	//printNode(listNode)
	return listNode
}

func printNode(h *ListNode) {
	for n := h; n != nil; n = n.Next {
		fmt.Println("listNode print val", n.Val)
	}
}

func reverse(head *ListNode) *ListNode {
	pre := head      //保存新的头节点
	cur := head.Next //保存当前节点
	head.Next = nil  //孤立头节点

	for cur != nil {
		next := cur.Next //保存下一个节点
		cur.Next = pre   //指向前节点
		pre = cur        //保存最新头节点
		cur = next       //重置当前节点
	}

	return pre
}

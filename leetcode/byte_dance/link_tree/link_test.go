package main

import (
	"fmt"
	"testing"
)

func Test_reverse(t *testing.T) {
	nodes := getNode()

	printNode(reverseList(nodes))

}

// convert *ListNode to []int
func L2s(head *ListNode) []int {
	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

// convert []int to *ListNode
func S2l(nums []int) *ListNode {
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

type question445 struct {
	para445
	ans445
}

// para 是参数
// one 代表第一个参数
type para445 struct {
	one     []int
	another []int
}

// ans 是答案
// one 代表第一个答案
type ans445 struct {
	one []int
}

func Test_addTwoNumbers(t *testing.T) {

	qs := []question445{

		question445{
			para445{[]int{}, []int{}},
			ans445{[]int{}},
		},

		question445{
			para445{[]int{1}, []int{1}},
			ans445{[]int{2}},
		},

		question445{
			para445{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
			ans445{[]int{2, 4, 6, 8}},
		},

		question445{
			para445{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
			ans445{[]int{2, 4, 6, 9, 0}},
		},

		question445{
			para445{[]int{1}, []int{9, 9, 9, 9, 9}},
			ans445{[]int{1, 0, 0, 0, 0, 0}},
		},

		question445{
			para445{[]int{9, 9, 9, 9, 9}, []int{1}},
			ans445{[]int{1, 0, 0, 0, 0, 0}},
		},

		question445{
			para445{[]int{2, 4, 3}, []int{5, 6, 4}},
			ans445{[]int{8, 0, 7}},
		},

		question445{
			para445{[]int{1, 8, 3}, []int{7, 1}},
			ans445{[]int{2, 5, 4}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 445------------------------\n")

	for _, q := range qs {
		_, p := q.ans445, q.para445
		fmt.Printf("【input】:%v       【output】:%v\n", p, L2s(addTwoNumbers(S2l(p.one), S2l(p.another))))
	}
	fmt.Printf("\n\n\n")
}

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
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

//反转一个单链表。
//
//示例:
//
//输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL
//进阶:
//你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
func reverseList(head *ListNode) *ListNode {
	if head==nil{
		return nil
	}

	var pre *ListNode
	pre = nil
	cur:=head

	for cur!=nil{
		next:=cur.Next
		cur.Next=pre
		pre=cur
		cur=next
	}
	return pre
}

//给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，
//并且它们的每个节点只能存储 一位 数字。
//
//如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
//
//您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

//示例：
//
//输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	l1Length := getLength(l1)
	l2Length := getLength(l2)

	newHeader := &ListNode{Val: 1, Next: nil}
	if l1Length < l2Length {
		newHeader.Next = addNode(l2, l1, l2Length-l1Length)
	} else {
		newHeader.Next = addNode(l1, l2, l1Length-l2Length)
	}

	if newHeader.Next.Val > 9 {
		newHeader.Next.Val = newHeader.Next.Val % 10
		return newHeader
	}

	return reverseList(newHeader.Next)
}


func addNode(l1 *ListNode, l2 *ListNode, offset int) *ListNode {
	if l1 == nil {
		return nil
	}

	var (
		res, node *ListNode
	)

	if offset == 0 {//两条链表长度一致
		res = &ListNode{Val: l1.Val + l2.Val, Next: nil}//首节点相加
		node = addNode(l1.Next, l2.Next, 0)
	} else {//两条链表长度不一致 l1>l2
		res = &ListNode{Val: l1.Val, Next: nil}
		node = addNode(l1.Next, l2, offset-1)
	}

	if node != nil && node.Val > 9 {
		res.Val++
		node.Val = node.Val % 10
	}
	res.Next = node
	return res
}

func getLength(l *ListNode) int {
	count := 0
	cur := l
	for cur != nil {
		count++
		cur = cur.Next
	}
	return count
}
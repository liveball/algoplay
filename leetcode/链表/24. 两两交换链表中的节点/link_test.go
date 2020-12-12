package link

import (
	"fmt"
	"testing"
)

func Test_24(t *testing.T) {

	nodes := getNode()
	//printNode(nodes)
	printNode(swapPairs(nodes))

	//printNode(swapPairs2(nodes))
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

func swapPairs2(head *ListNode) *ListNode { //非递归交换

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

// 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
// 示例 1：
// 输入：head = [1,2,3,4]
// 输出：[2,1,4,3]
// 示例 2：

// 输入：head = []
// 输出：[]
// 示例 3：

// 输入：head = [1]
// 输出：[1]
// 提示：

// 链表中节点的数目在范围 [0, 100] 内
// 0 <= Node.val <= 100

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/swap-nodes-in-pairs
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

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

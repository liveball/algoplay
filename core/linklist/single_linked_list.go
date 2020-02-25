package main

import (
	"fmt"
)

type Node struct {
	Val  int
	Next *Node
}

//打印链表
func Print(head *Node) {
	cur := head
	format := ""

	for nil != cur {
		format += fmt.Sprintf("%+v", cur.Val)
		cur = cur.Next

		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}

//单向循环链表
func CycleLinkedList(head *Node) *Node {
	var t *Node

	for h := head; h != nil; h = h.Next {
		if h.Next == nil {
			t = h
		}
		//fmt.Println(h)

	}

	t.Next = head
	return head
}

// 判断链表是否有环，不能使用额外的空间。
// 给 2 个指针，一个指针是另外一个指针的下一个指针。
// 快指针一次走 2 格，慢指针一次走 1 格。
// 如果存在环，那么前一个指针一定会经过若干圈之后追上慢的指针。

//判断单链表是否为单向循环链表
func HasCycle(head *Node) bool {
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

//单链表反转
func ReverseList(head *Node) *Node {
	var pre *Node
	pre = nil
	cur := head

	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}

//在某个节点后面插入节点
func InsertAfter(p *Node, val int) *Node {
	if p == nil {
		return nil
	}

	newNode := &Node{Val: val}
	oldNext := p.Next
	p.Next = newNode
	newNode.Next = oldNext

	return p
}

//在链表头部插入节点
func InsertToHead(head *Node, val int) *Node {
	return InsertAfter(head, val)
}

//在链表尾部插入节点
func InsertToTail(head *Node, val int) *Node {
	cur := head
	for nil != cur.Next {
		cur = cur.Next
	}
	return InsertAfter(cur, val)
}

//在某个节点前面插入节点
func InsertBefore(head, p *Node, val int) *Node {
	if p == nil || p == head {
		return nil
	}

	cur := head.Next
	pre := head
	for cur != nil {
		if cur == p {
			break
		}

		pre = cur
		cur = cur.Next
	}

	if cur == nil {
		return nil
	}

	newNode := &Node{Val: val}
	pre.Next = newNode
	newNode.Next = cur

	return head
}

//删除某个值的节点
/**
 * 方法二：时间复杂度O(N),额外空间复杂度O(1)
 * 不需要任何容器，只需要几个有限的变量直接调整原链表
 * */
func DelNodeByVal(head *Node, val int) *Node {
	//找到第一个不等于val的值，该节点一定不会被删除，而且必是调整后的链表的头结点
	for head != nil {
		if head.Val != val {
			break
		}
		head = head.Next
		//Print(head)
	}

	//循环结束时head指向调整后的链表的头结点
	//定义几个有限的变量完成链表的调整
	cur := head
	pre := head
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
		} else {
			pre = cur
		}

		cur = cur.Next
	}

	//调整结束返回新链表的头结点
	return head
}

//删除传入的节点
func DelNode(head *Node, p *Node) *Node {
	if head == nil || p == nil {
		return nil
	}

	cur := head.Next
	pre := head
	for cur != nil {
		if cur == p {
			break
		}

		pre = cur
		cur = cur.Next
	}

	if cur == nil {
		return nil
	}

	pre.Next = p.Next
	p = nil

	return head
}

//删除倒数第n个节点
func DelBottomN(head *Node, n int) *Node {
	if n <= 0 || head == nil {
		return nil
	}

	fast := head
	for i := 1; i <= n && fast != nil; i++ {
		fast = fast.Next
	}

	if fast == nil {
		return nil
	}

	slow := head
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next

	return head
}

//通过索引查找节点
func FindByIndex(head *Node, index int) *Node {
	if head == nil {
		return nil
	}

	for i := 0; i < index; i++ {
		head = head.Next
	}

	return head
}

//单链表交换
func SwapList(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	var dummy Node
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

//两个有序链表合并
func MergeSortedList(l1, l2 *Node) *Node {
	if l1 == nil || l1.Next == nil {
		return l2
	}

	if l2 == nil || l2.Next == nil {
		return l1
	}

	head := &Node{}

	if l1.Val > l2.Val {
		head = l2
	} else {
		head = l1
	}

	cur := head

	curl1 := l1.Next
	curl2 := l2.Next

	for curl1 != nil && curl2 != nil {
		if curl1.Val > curl2.Val {
			cur.Next = curl2
			curl2 = curl2.Next

		} else {
			cur.Next = curl1
			curl1 = curl1.Next
		}

		cur = cur.Next
	}

	if curl1 != nil {
		cur.Next = curl1
	} else if curl2 != nil {
		cur.Next = curl2
	}

	return head
}

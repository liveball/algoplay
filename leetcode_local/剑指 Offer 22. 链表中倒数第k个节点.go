package main

//  初始化： 前指针 former 、后指针 latter ，双指针都指向头节点 head​ 。
// 构建双指针距离： 前指针 former 先向前走 k 步（结束后，双指针 former 和 latter 间相距 kk 步）。
// 双指针共同移动： 循环中，双指针 former 和 latter 每轮都向前走一步，直至 former 走过链表 尾节点 时跳出（跳出后， latter 与尾节点距离为 k-1，即 latter 指向倒数第 k个节点）。
// 返回值： 返回 latter 即可。

/* Definition for singly-linked list.*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	slow, fast := head, head
	for i := 0; i < k; i++ {
		if fast != nil {
			fast = fast.Next
		}
	}

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}

func getKthFromEnd2(head *ListNode, k int) *ListNode {
	slow, fast := head, head

	t := 0
	for fast != nil {
		fast = fast.Next
		t++
		if t > k {
			slow = slow.Next
		}
	}

	return slow
}

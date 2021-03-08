package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	l1 := headA
	l2 := headB
	for l1 != l2 {
		if l1 != nil {
			l1 = l1.Next
		} else {
			l1 = headB
		}

		if l2 != nil {
			l2 = l2.Next
		} else {
			l2 = headA
		}
	}

	return l1
}

// l1: 1 3
// l2: 2 4 nil
func getIntersectionNode2(l1, l2 *ListNode) (res *ListNode) {
	if l1 == nil || l2 == nil {
		res = nil
		return
	}

	tmp := l2
	for l1 != nil {
		l2 := tmp
		for l2 != nil {
			if l1.Val == l2.Val {
				res = &ListNode{
					Val: l1.Val,
				}
			}
			l2 = l2.Next
		}
		l1 = l1.Next
	}

	return
}

func getIntersectionNode3(l1, l2 *ListNode) (res *ListNode) {
	if l1 == nil || l2 == nil {
		res = nil
		return
	}

	tmpMap := make(map[*ListNode]*ListNode)
	for l1 != nil {
		tmpMap[l1] = l1
		l1 = l1.Next
	}
	for l2 != nil {
		if v, ok := tmpMap[l2]; ok {
			res = v
			return
		}
		l2 = l2.Next
	}

	return
}

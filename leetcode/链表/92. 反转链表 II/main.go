package main

func main(){
	
}

/**
 * Definition for singly-linked list.
 */
 type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
    if head == nil || m > n {
        return nil
    }
    
    //申请节点指向头节点
    dummy := &ListNode{}
    dummy.Next = head
    prve := dummy
    
   //走到将要翻转节点的前一个节点 prev
    for i := 0; i < m-1; i++ {
        prve = prve.Next
    }

    //cur 第m个节点，也就是将要翻转的节点
    cur := prve.Next


    for i := m; i < n; i++ {
        next := cur.Next        //保存要反转节点的下一个节点
        cur.Next = next.Next    //当前节点指向 要放转节点的next节点，最终指向第m个节点的next
        next.Next = prve.Next  //第n个节点的next指向前一个节点
        prve.Next = next       // 第m个节点指向后面一个节点
    }
    
    return dummy.Next    
}

// 92. 反转链表 II
// 反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。

// 说明:
// 1 ≤ m ≤ n ≤ 链表长度。

// 示例:

// 输入: 1->2->3->4->5->NULL, m = 2, n = 4
// 输出: 1->4->3->2->5->NULL
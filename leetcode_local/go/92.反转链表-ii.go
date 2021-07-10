/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	   // base case
	   if (left == 1) {
        // 相当于反转前 right个元素
        return reverseN(head, right);
       }

	   // 前进到反转的起点触发 base case
	   head.Next=reverseBetween(head.Next, left-1, right-1)
	   return head
}

var successor  *ListNode//后驱节点

func reverseN(head *ListNode, n int) *ListNode {
   if n==1{
	  successor=head.Next
	  return head
   }
    // 以 head.Next 为起点，需要反转前 n - 1 个节点
	last:=reverseN(head.Next, n-1)
	head.Next.Next=head
	// 让反转之后的 head 节点和后面的节点连起来
	head.Next=successor
    return last
}
// @lc code=end


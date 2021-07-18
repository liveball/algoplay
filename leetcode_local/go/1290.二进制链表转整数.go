/*
 * @lc app=leetcode.cn id=1290 lang=golang
 *
 * [1290] 二进制链表转整数
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func getDecimalValue(head *ListNode) int {
	sum := 0
	for head != nil {
		sum = sum*2 + head.Val
		head = head.Next
	}
	return sum
}

// @lc code=end


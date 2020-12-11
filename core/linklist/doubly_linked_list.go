package main

// 链表的一个节点
type ListNode struct {
	prev  *ListNode // 前一个节点
	next  *ListNode // 后一个节点
	value string    // 数据
}

// 创建一个节点
func NewListNode(value string) (listNode *ListNode) {
	listNode = &ListNode{
		value: value,
	}

	return
}

// 当前节点的前一个节点
func (n *ListNode) Prev() (prev *ListNode) {
	prev = n.prev

	return
}

// 当前节点的前一个节点
func (n *ListNode) Next() (next *ListNode) {
	next = n.next

	return
}

// 获取节点的值
func (n *ListNode) GetValue() (value string) {
	if n == nil {

		return
	}
	value = n.value

	return
}

package main

// 链表
type List struct {
	head *ListNode // 表头节点
	tail *ListNode // 表尾节点
	len  int       // 链表的长度
}

// 创建一个空链表
func NewList() (list *List) {
	list = &List{}
	return
}

// 返回链表头节点
func (l *List) Head() (head *ListNode) {
	head = l.head

	return
}

// 返回链表尾节点
func (l *List) Tail() (tail *ListNode) {
	tail = l.tail

	return
}

// 返回链表长度
func (l *List) Len() (len int) {
	len = l.len

	return
}

// 在链表的右边插入一个元素
func (l *List) RPush(value string) {

	node := NewListNode(value)
	// 链表为空的时候
	if l.Len() == 0 {
		l.head = node
		l.tail = node
	} else {
		tail := l.tail
		tail.next = node
		node.prev = tail

		l.tail = node
	}

	l.len++

	return
}

// 从链表左边取出一个节点
func (l *List) LPop() (node *ListNode) {

	// 数据为空
	if l.len == 0 {

		return
	}

	node = l.head

	if node.next == nil { // 链表为空
		l.head = nil
		l.tail = nil
	} else {
		l.head = node.next
	}

	l.len--
	return
}

// 通过索引查找节点
// 查不到节点则返回空
func (l *List) Index(index int) (node *ListNode) {

	// 索引为负数则表尾开始查找
	if index < 0 {
		index = (-index) - 1
		node = l.tail
		for true {
			// 未找到
			if node == nil {

				return
			}

			// 查到数据
			if index == 0 {

				return
			}

			node = node.prev
			index--
		}
	} else {
		node = l.head
		for ; index > 0 && node != nil; index-- {
			node = node.next
		}
	}

	return
}

// 返回指定区间的元素
func (l *List) Range(start, stop int) (nodes []*ListNode) {
	nodes = make([]*ListNode, 0)

	// 转为自然数
	if start < 0 {
		start = l.len + start
		if start < 0 {
			start = 0
		}
	}

	if stop < 0 {
		stop = l.len + stop
		if stop < 0 {
			stop = 0
		}
	}

	// 区间个数
	rangeLen := stop - start + 1
	if rangeLen < 0 {

		return
	}

	startNode := l.Index(start)
	for i := 0; i < rangeLen; i++ {
		if startNode == nil {
			break
		}

		nodes = append(nodes, startNode)
		startNode = startNode.next
	}

	return
}

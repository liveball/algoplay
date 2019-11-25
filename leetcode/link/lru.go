package link

type Node struct {
	Next *Node
	Val  int
}

type LRU struct {
	Head   *Node
	Tail   *Node
	Length int
	Cap    int
}

// 用链表实现lru 缓存队列

//1、如果缓存队列中命中该数据
//     将该数据对应的节点删除，并将该数据添加到链表的头部

//2、如果缓存队列中没有命中该数据
//     若缓存未满则直接将数据添加到链表头部
//     若缓存已满，则先删除链表尾部的节点，再将该数据添加到链表头部

func New() (lru *LRU) {
	lru = &LRU{
		Cap: 10,
	}
	return
}

func (lru *LRU) Init() {
	lru.Cap = 0    // 此时链表是空的
	lru.Head = nil // 没有车头
	lru.Tail = nil // 没有车尾
}

func (lru *LRU) Append2(node *Node) {
	if lru.Length == 0 { // 无元素的时候添加
		lru.Head = node // 这是单链表的第一个元素，也是链表的头部
		lru.Tail = node // 同时是单链表的尾部
		lru.Length = 1  // 单链表有了第一个元素
	} else { // 有元素了再添加
		oldTail := lru.Tail
		oldTail.Next = node // node放到尾部元素后面
		lru.Tail = node     // node成为新的尾部
		lru.Length++        // 元素数量增加
	}
}

func (lru *LRU) Append(node *Node) bool {
	if node == nil {
		return false
	}

	node.Next = nil
	// 将新元素放入单链表中
	if lru.Length == 0 {
		lru.Head = node
	} else {
		oldTail := lru.Tail
		oldTail.Next = node
	}

	// 调整尾部位置，及链表元素数量
	lru.Tail = node // node成为新的尾部
	lru.Length++    // 元素数量增加

	return true
}

func (lru *LRU) Query(key int) (hit bool) {
	if lru.Node == nil {
		hit = false
		return
	}

	head := lru.Node
	for n := head; n != nil; n = n.Next {
		if n.Val == key {
			hit = true
		}

		lru.Length++ //计算队列长度
	}

	return
}

func (lru *LRU) AddCache(key int) {
	if lru.Query(key) { //缓存队列中存在

	} else { //缓存队列中不存在
		if lru.Length > lru.Cap { //超过缓存容量

		} else {

		}
	}
}

func (lru *LRU) AddNode(key int) {
	if lru.Length == 0 {
		lru.Node = &Node{
			Val: key,
		}
		return
	}

	if lru.Length > 0 {
		cur := lru.Node
		head := &Node{
			Val: key,
		}
		head.Next = cur
		lru.Node = head
	}
}

func (lru *LRU) DelNode(key int) {
	if lru.Length == 0 {
		return
	}

	if lru.Length > 0 {
		for n := lru.Node; n != nil; n = n.Next {
			if n.Val == key {
				n = n.Next
				break
			}
		}
	}
}

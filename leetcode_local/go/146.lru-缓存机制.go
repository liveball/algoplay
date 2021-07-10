/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存机制
 */

// @lc code=start
type LRUCache struct {
	cacheMap       map[int]*Node
	doubleLinkList *LinkList
	capacity       int
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cacheMap:       make(map[int]*Node, capacity),
		doubleLinkList: NewLinkList(),
		capacity:       capacity,
	}

	return l
}

/* 将某个 key 提升为最近使用的 */
func (this *LRUCache) makeRecently(key int) {
	x, ok := this.cacheMap[key]
	if !ok {
		return
	}
	// 先从链表中删除这个节点
	this.doubleLinkList.Remove(x)
	// 重新插到队尾
	this.doubleLinkList.AddNodeToTail(x)
}

/* 添加最近使用的元素 */
func (this *LRUCache) addRecently(key, val int) {
	x := &Node{
		key: key,
		val: val,
	}
	// 链表尾部就是最近使用的元素
	this.doubleLinkList.AddNodeToTail(x)
	// 别忘了在 map 中添加 key 的映射
	this.cacheMap[key] = x
}

/* 删除某一个 key */
func (this *LRUCache) deleteKey(key int) {
	x, ok := this.cacheMap[key]
	if !ok {
		return
	}
	// 从链表中删除
	this.doubleLinkList.Remove(x)
	// 从 map 中删除
	delete(this.cacheMap, key)
}

/* 删除最久未使用的元素 */
func (this *LRUCache) removeLeastRecently() {
	// 链表头部的第一个元素就是最久未使用的
	deletedNode := this.doubleLinkList.RemoveFirst()
	if deletedNode == nil {
		return
	}

	// 同时别忘了从 map 中删除它的 key
	delete(this.cacheMap, deletedNode.key)
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.cacheMap[key]
	if !ok {
		return -1
	}
	// 将该数据提升为最近使用的
	this.makeRecently(key)

	if v == nil {
		return -1
	}

	return v.val
}

/* 添加最近使用的元素 */
func (this *LRUCache) Put(key int, value int) {
	_, ok := this.cacheMap[key]
	if ok {
		// 删除旧的数据
		this.deleteKey(key)
		// 新插入的数据为最近使用的数据
		this.addRecently(key, value)
		return
	}

	if this.capacity == this.doubleLinkList.Size() {
		// 删除最久未使用的元素
		this.removeLeastRecently()
	}

	// 添加为最近使用的元素
	this.addRecently(key, value)
}

type Node struct {
	key  int
	val  int
	prev *Node
	next *Node
}

type LinkList struct {
	//链表元素个数
	size int
	//头尾虚节点
	head *Node
	tail *Node
}

func NewLinkList() *LinkList {
	l := &LinkList{
		head: &Node{},
		tail: &Node{},
	}

	l.head.next = l.tail
	l.tail.prev = l.head

	return l
}

// 在链表尾部添加节点 x，时间 O(1)
func (l *LinkList) AddNodeToTail(x *Node) {
	x.prev = l.tail.prev
	x.next = l.tail
	l.tail.prev.next = x
	l.tail.prev = x
	l.size++
}

// 删除链表中的 x 节点（x 一定存在）
// 由于是双链表且给的是目标 Node 节点，时间 O(1)
func (l *LinkList) Remove(x *Node) {
	x.prev.next = x.next
	x.next.prev = x.prev
	l.size--
}

//删除链表中第一个节点，并返回该节点，时间 O(1)
func (l *LinkList) RemoveFirst() *Node {
	if l.head.next == l.tail { //头尾节点相遇，无节点返回
		return nil
	}

	first := l.head.next
	l.Remove(first)
	return first
}

// 返回链表长度，时间 O(1)
func (l *LinkList) Size() int {
	return l.size
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end


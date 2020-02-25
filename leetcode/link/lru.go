package link

type Node struct {
	Key  int
	Val  int
	Next *Node
	Prev *Node
}

type LRUCache struct {
	limit int
	hash  map[int]*Node
	Head  *Node
	Tail  *Node
}

func Constructor(capacity int) LRUCache {
	h := &Node{-1, -1, nil, nil}
	t := &Node{-1, -1, nil, nil}
	h.Next = t
	t.Prev = h

	hash := make(map[int]*Node, capacity)
	return LRUCache{
		hash:  hash,
		limit: capacity,
		Head:  h,
		Tail:  t,
	}
}

func (this *LRUCache) insert(node *Node) {
	t := this.Tail
	node.Prev = t.Prev
	t.Prev.Next = node
	node.Next = t
	t.Prev = node
}

func (this *LRUCache) remove(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.hash[key]; ok {
		this.remove(v)
		this.insert(v)
		return v.Val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key, val int) {
	if v, ok := this.hash[key]; ok {
		this.remove(v)
		this.insert(v)
		v.Val = val
	} else {
		if len(this.hash) >= this.limit {//如果此时缓存已满，则链表尾结点删除，将新的数据结点插入链表的头部
			h := this.Head.Next
			this.remove(h)
			delete(this.hash, h.Key)
		}

		//如果此时缓存未满，则将此结点直接插入到链表的头部
		node := &Node{key, val, nil, nil}
		this.hash[key] = node
		this.insert(node)
	}
}


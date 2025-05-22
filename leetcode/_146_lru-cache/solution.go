package _146_lru_cache

type LRUCache struct {
	m        map[int]*DLinkedNode
	head     *DLinkedNode
	tail     *DLinkedNode
	capacity int
}

type DLinkedNode struct {
	Key  int
	Val  int
	Pre  *DLinkedNode
	Next *DLinkedNode
}

func Constructor(capacity int) LRUCache {
	head := &DLinkedNode{Val: -1}
	tail := &DLinkedNode{Val: -1, Pre: head}
	head.Next = tail
	return LRUCache{
		m:        make(map[int]*DLinkedNode),
		head:     head,
		tail:     tail,
		capacity: capacity,
	}
}

func (l *LRUCache) Get(key int) int {
	if node, ok := l.m[key]; ok {
		l.MoveToHead(node)
		return node.Val
	}

	return -1
}

func (l *LRUCache) Put(key int, value int) {
	if node, ok := l.m[key]; ok {
		l.MoveToHead(node)
		node.Val = value
		return
	}

	if len(l.m) >= l.capacity {
		node := l.RemoveTail()
		delete(l.m, node.Key)
	}

	node := &DLinkedNode{Key: key, Val: value}
	l.AddToHead(node)
	l.m[key] = node
}

// A <-> B
// A <-> C <-> B
func (l *LRUCache) MoveToHead(node *DLinkedNode) {
	// 从现有位置移除
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre

	// 把节点插入到l.head之后
	l.AddToHead(node)
}

func (l *LRUCache) AddToHead(node *DLinkedNode) {
	node.Next = l.head.Next
	l.head.Next.Pre = node
	node.Pre = l.head
	l.head.Next = node
}

func (l *LRUCache) RemoveTail() *DLinkedNode {
	target := l.tail.Pre
	target.Pre.Next = target.Next
	target.Next.Pre = target.Pre
	return target
}

//func (l *LRUCache) Info() {
//	sb := strings.Builder{}
//	sb.WriteString("map:\n")
//	for k, v := range l.m {
//		sb.WriteString(fmt.Sprintf("[%v, %v],", k, v.Val))
//	}
//	sb.WriteString("\n")
//
//	sb.WriteString(fmt.Sprintf("list head:%v, list tail:%v\n", l.head.Next, l.tail))
//	sb.WriteString("list detail:\n")
//	p := l.head.Next
//	for p != nil {
//		sb.WriteString(fmt.Sprintf("_%v_", p.Val))
//		p = p.Next
//	}
//
//	println(sb.String())
//}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

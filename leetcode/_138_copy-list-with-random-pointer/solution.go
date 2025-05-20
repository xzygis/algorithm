package _138_copy_list_with_random_pointer

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}

	// 复制
	for p := head; p != nil; p = p.Next.Next {
		p.Next = &Node{Val: p.Val, Next: p.Next}
	}

	// 赋值复制后的节点的random指针
	for p := head; p != nil; p = p.Next.Next {
		if p.Random != nil {
			p.Next.Random = p.Random.Next
		}

	}

	// 拆开链表
	newHead := head.Next
	for p1 := head; p1 != nil; p1 = p1.Next {
		p2 := p1.Next
		p1.Next = p1.Next.Next
		if p2.Next != nil {
			p2.Next = p2.Next.Next
		}
	}

	return newHead
}

func copyRandomListV1(head *Node) *Node {
	nodeMap := make(map[*Node]*Node)
	p := head
	for p != nil {
		nodeMap[p] = &Node{Val: p.Val}
		p = p.Next
	}

	p = head
	for p != nil {
		nodeMap[p].Next = nodeMap[p.Next]
		nodeMap[p].Random = nodeMap[p.Random]
		p = p.Next
	}

	return nodeMap[head]
}

package _35_fu_za_lian_biao_de_fu_zhi_lcof

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

type Node struct {
	Val int
	Next *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	//复制链表节点，并插入原链表
	p := head
	for p != nil {
		cloned := &Node{Val: p.Val}
		cloned.Next = p.Next
		p.Next = cloned
		p = cloned.Next
	}

	//给复制出来的节点赋值Rand指针
	p = head
	for p != nil {
		if p.Random != nil {
			p.Next.Random = p.Random.Next
		}
		p = p.Next.Next
	}

	//拆分链表
	newHead := head.Next
	p = head
	p1 := newHead
	for p != nil {
		p.Next = p1.Next
		if p1.Next != nil {
			p1.Next = p1.Next.Next
		}
		p = p.Next
		p1 = p1.Next
	}

	return newHead
}

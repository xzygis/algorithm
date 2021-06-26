package _35_clone_complex_list

/*
面试题35：复杂链表的赋值
 */

import (
	"fmt"
	"math/rand"
)

type ComplexListNode struct {
	val     int
	next    *ComplexListNode
	sibling *ComplexListNode
}

// Clone 以空间换时间，保存<N, N'>的映射关系
func Clone(head *ComplexListNode) *ComplexListNode {
	if head == nil {
		return nil
	}

	m := make(map[*ComplexListNode]*ComplexListNode)
	p := head
	for p != nil {
		m[p] = &ComplexListNode{val: p.val}
		p = p.next
	}

	p = head
	for p != nil {
		m[p].next = m[p.next]
		m[p].sibling = m[p.sibling]
		p = p.next
	}

	return m[head]
}

// CloneV2 先复制链表，然后原地保存，再赋值兄弟指针，最后拆分链表
func CloneV2(head *ComplexListNode) *ComplexListNode {
	if head == nil {
		return nil
	}

	// 拷贝原有节点，并串联到链表中
	cur := head
	for cur != nil {
		next := cur.next
		newP := &ComplexListNode{val: cur.val}
		cur.next = newP
		newP.next = next
		cur = next
	}

	// 赋值sibling
	cur = head
	for cur != nil {
		if cur.sibling != nil {
			cur.next.sibling = cur.sibling.next
		}
		cur = cur.next.next
	}

	// 拆分新老链表
	newHead := head.next
	cur = head
	clonedCur := cur.next
	for cur.next.next != nil {
		cur.next = cur.next.next
		clonedCur.next = cur.next.next
		cur = cur.next
		clonedCur = clonedCur.next
	}
	cur.next = nil

	return newHead
}

func createComplexList(len int) *ComplexListNode {
	var head *ComplexListNode
	var pre *ComplexListNode
	for i := 0; i < len; i++ {
		if head == nil {
			head = &ComplexListNode{val: i + 1}
			pre = head
		} else {
			pre.next = &ComplexListNode{val: i + 1}
			pre = pre.next
		}
	}

	p := head
	for i := 0; i < len; i++ {
		step := int(rand.Uint32()) % len
		tmp := head
		for j := 0; j < step; j++ {
			tmp = tmp.next
		}
		p.sibling = tmp
		p = p.next
	}
	return head
}

func printComplexList(head *ComplexListNode) {
	for head != nil {
		printComplexListNode(head)
		head = head.next
	}
	fmt.Println()
}

func printComplexListNode(node *ComplexListNode) {
	if node.next != nil && node.sibling != nil {
		fmt.Printf("Node(%d,%d,%d)", node.val, node.next.val, node.sibling.val)
	} else if node.next != nil {
		fmt.Printf("Node(%d,%d,-)", node.val, node.next.val)
	} else if node.sibling != nil {
		fmt.Printf("Node(%d,-,%d)", node.val, node.sibling.val)
	} else {
		fmt.Printf("Node(%d,-,-)", node.val)
	}
}
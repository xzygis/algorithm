package _138_copy_list_with_random_pointer

import (
	"fmt"
	"strings"
	"testing"
)

func Test_copyRandomList(t *testing.T) {
	type args struct {
		head *Node
	}

	node7 := &Node{Val: 7}
	node13 := &Node{Val: 13}
	node11 := &Node{Val: 11}
	node10 := &Node{Val: 10}
	node1 := &Node{Val: 1}
	node7.Next = node13
	node13.Next = node11
	node11.Next = node10
	node10.Next = node1
	node13.Random = node7
	node11.Random = node1
	node10.Random = node11
	node1.Random = node7

	tests := []struct {
		name string
		args args
		want *Node
	}{
		{

			name: "test case 1",
			args: args{
				head: node7,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := copyRandomList(tt.args.head)
			t.Log(got)
			PrintNodes(got)
		})
	}
}

func PrintNodes(node *Node) {
	sb := strings.Builder{}
	for node != nil {
		str := "_["
		str += fmt.Sprintf("%v", node.Val)
		if node.Next != nil {
			str += fmt.Sprintf("_next_%v", node.Next.Val)
		}
		if node.Random != nil {
			str += fmt.Sprintf("_rand_%v", node.Random.Val)
		}
		str += "]"

		sb.WriteString(str)
		node = node.Next
	}
	println(sb.String())
}

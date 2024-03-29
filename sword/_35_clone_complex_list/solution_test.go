package _35_clone_complex_list


import (
	"fmt"
	"testing"
)

func TestClone(t *testing.T) {
	type args struct {
		head *ComplexListNode
	}
	tests := []struct {
		name string
		args args
		want *ComplexListNode
	}{
		{
			name: "test case 1",
			args: args{
				head: createComplexList(10),
			},
			want: createComplexList(10),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printComplexList(tt.args.head)
			got := Clone(tt.args.head)
			printComplexList(got)
			assert(t, tt.args.head, got)

			printComplexList(tt.args.head)
			got = CloneV2(tt.args.head)
			printComplexList(got)
			assert(t, tt.args.head, got)
		})
	}
}

func assert(t *testing.T, origin *ComplexListNode, got *ComplexListNode) {
	head := origin
	newHead := got
	for head != nil {
		if head.val != newHead.val {
			t.Fatal()
		}

		if head.next != nil {
			if head.next.val != newHead.next.val {
				t.Fatal()
			}
		}

		if head.sibling != nil {
			if head.sibling.val != newHead.sibling.val {
				t.Fatal()
			}
		}
		head = head.next
		newHead = newHead.next
	}
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

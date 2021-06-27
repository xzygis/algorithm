package _35_fu_za_lian_biao_de_fu_zhi_lcof

import (
	"reflect"
	"testing"
)

func Test_copyRandomList(t *testing.T) {
	type args struct {
		head *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "test case 1",
			args: args{
				head: createList(),
			},
			want: createList(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := copyRandomList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				old := tt.args.head
				p := old
				p1 := got
				for p != nil {
					if p.Val != p1.Val {
						t.Fatal()
					}

					if p.Next != nil {
						if p.Next.Val != p1.Next.Val {
							t.Fatal()
						}
					}

					if p.Random != nil {
						if p.Random.Val != p1.Random.Val {
							t.Fatal(p1.Random)
						}
					}

					p = p.Next
					p1 = p1.Next
				}
			}
		})
	}
}

//[[7,null],[13,0],[11,4],[10,2],[1,0]]
func createList() *Node {
	node0 := &Node{Val: 7}
	node1 := &Node{Val: 13}
	node2 := &Node{Val: 11}
	node3 := &Node{Val: 10}
	node4 := &Node{Val: 1}

	node0.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	node1.Random = node0
	node2.Random = node4
	node3.Random = node2
	node4.Random = node0

	return node0
}

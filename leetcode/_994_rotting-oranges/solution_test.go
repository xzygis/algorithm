package _994_rotting_oranges

import (
	"testing"

	"github.com/xzygis/algorithm/leetcode/utils"
)

func Test_orangesRotting(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 1",
			args: args{
				matrix: utils.ParseMatrix("[[2,1,1],[1,1,0],[0,1,1]]"),
			},
			want: 4,
		},
		{
			name: "test case 2",
			args: args{matrix: utils.ParseMatrix("[[2,1,1],[0,1,1],[1,0,1]]")},
			want: -1,
		},
		{
			name: "test case 3",
			args: args{matrix: utils.ParseMatrix("[[0,2]]")},
			want: 0,
		},
		{
			name: "test case 4",
			args: args{matrix: utils.ParseMatrix("[[1],[2]]")},
			want: 1,
		},
		{
			name: "test case 5",
			args: args{matrix: utils.ParseMatrix("[[1,2]]")},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := orangesRotting(tt.args.matrix); got != tt.want {
				t.Errorf("orangesRotting() = %v, want %v", got, tt.want)
			}
		})
	}
}

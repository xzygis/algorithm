package _207_course_schedule

import (
	"testing"

	"github.com/xzygis/algorithm/leetcode/utils"
)

func Test_canFinish(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test case 1",
			args: args{
				numCourses:    2,
				prerequisites: utils.ParseMatrix("[[1,0]]"),
			},
			want: true,
		},
		{
			name: "test case 2",
			args: args{
				numCourses:    2,
				prerequisites: utils.ParseMatrix("[[1,0],[0,1]]"),
			},
			want: false,
		},
		{
			name: "test case 3",
			args: args{
				numCourses:    3,
				prerequisites: utils.ParseMatrix("[[1,0],[0,2],[2,1]]"),
			},
			want: false,
		},
		{
			name: "test case 4",
			args: args{
				numCourses:    3,
				prerequisites: utils.ParseMatrix("[[1,0],[2,0],[0,2]]"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canFinish(tt.args.numCourses, tt.args.prerequisites); got != tt.want {
				t.Errorf("canFinish() = %v, want %v", got, tt.want)
			}
		})
	}
}

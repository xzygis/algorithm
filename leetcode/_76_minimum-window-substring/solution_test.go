package _76_minimum_window_substring

import (
	"testing"
)

func Test_minWindow(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test case 1",
			args: args{
				s: "ADOBECODEBANC",
				t: "ABC",
			},
			want: "BANC",
		},
		{
			name: "test case 2",
			args: args{
				s: "a",
				t: "a",
			},
			want: "a",
		},
		{
			name: "test case 3",
			args: args{
				s: "abc",
				t: "abc",
			},
			want: "abc",
		},
		{
			name: "test case 4",
			args: args{
				s: "a",
				t: "aa",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindow(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("minWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

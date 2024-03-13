package longestpalindromicsubstring

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test case 1",
			args: args{
				s: "abba",
			},
			want: "abba",
		},
		{
			name: "test case 2",
			args: args{
				s: "babad",
			},
			want: "bab",
		},
		{
			name: "test case 3",
			args: args{
				s: "cbbd",
			},
			want: "bb",
		},
		{
			name: "test case 4",
			args: args{
				s: "",
			},
			want: "",
		},
		{
			name: "test case 5",
			args: args{
				s: "aaaaaaaaaa",
			},
			want: "aaaaaaaaaa",
		},
		{
			name: "test case 6",
			args: args{
				s: "aaaaaaaaaa888",
			},
			want: "aaaaaaaaaa",
		},
		{
			name: "test case 7",
			args: args{
				s: "aa",
			},
			want: "aa",
		},
		{
			name: "test case 7",
			args: args{
				s: "a",
			},
			want: "a",
		},
		{
			name: "test case 8",
			args: args{
				s: "ac",
			},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

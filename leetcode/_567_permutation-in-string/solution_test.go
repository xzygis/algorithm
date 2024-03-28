package permutationinstring

import "testing"

func Test_checkInclusion(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test case 1",
			args: args{
				s1: "ab",
				s2: "eidbaooo",
			},
			want: true,
		},
		{
			name: "test case 2",
			args: args{
				s1: "ab",
				s2: "eidboaoo",
			},
			want: false,
		},
		{
			name: "test case 3",
			args: args{
				s1: "abcd",
				s2: "cwabcaddef",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkInclusion(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("checkInclusion() = %v, want %v", got, tt.want)
			}
		})
	}
}

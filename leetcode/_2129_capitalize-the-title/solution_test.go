package _2129_capitalize_the_title

import (
	"testing"
)

func Test_capitalizeTitle(t *testing.T) {
	type args struct {
		title string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test case 1",
			args: args{
				title: "capiTalIze tHe titLe",
			},
			want: "Capitalize The Title",
		},

		{
			name: "test case 2",
			args: args{
				title: "capiTalIze OF titLe",
			},
			want: "Capitalize of Title",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := capitalizeTitle(tt.args.title); got != tt.want {
				t.Errorf("capitalizeTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

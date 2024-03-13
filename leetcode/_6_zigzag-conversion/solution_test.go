package zigzagconversion

import "testing"

func Test_convert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test case 1",
			args: args{
				s:       "PAYPALISHIRING",
				numRows: 3,
			},
			want: "PAHNAPLSIIGYIR",
		},
		{
			name: "test case 2",
			args: args{
				s:       "PAYPALISHIRING",
				numRows: 4,
			},
			want: "PINALSIGYAHRPI",
		},
		{
			name: "test case 3",
			args: args{
				s:       "A",
				numRows: 1,
			},
			want: "A",
		},
		{
			name: "test case 4",
			args: args{
				s:       "",
				numRows: 1,
			},
			want: "",
		},
		{
			name: "test case 5",
			args: args{
				s:       "AB",
				numRows: 1,
			},
			want: "AB",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}

package _221_maximal_square

import (
	"github.com/xzygis/algorithm/leetcode/utils"
	"testing"
)

func Test_maximalSquare(t *testing.T) {
	type args struct {
		matrix [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 1",
			args: args{
				matrix: MatrixIntToByte(utils.ParseMatrixStr(`[["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]`)),
			},
			want: 4,
		},
		{
			name: "test case 2",
			args: args{
				matrix: MatrixIntToByte(utils.ParseMatrixStr(`[["0","1"],["1","0"]]`)),
			},
			want: 1,
		},
		{
			name: "test case 3",
			args: args{
				matrix: MatrixIntToByte(utils.ParseMatrixStr(`[["0"]]`)),
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalSquare(tt.args.matrix); got != tt.want {
				t.Errorf("maximalSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func MatrixIntToByte(matrix [][]int) [][]byte {
	result := make([][]byte, len(matrix))
	if len(matrix) == 0 {
		return result
	}

	rows := len(matrix)
	cols := len(matrix[0])
	for i := 0; i < rows; i++ {
		result[i] = make([]byte, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result[i][j] = byte(matrix[i][j] + '0')
		}
	}
	return result
}

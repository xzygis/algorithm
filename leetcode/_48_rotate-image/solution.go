package _48_rotate_image

// [i][j] => [j][col-i-1]
// [i][j] = [col-j-1][i]
func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[n-j-1][i]
			matrix[n-j-1][i] = matrix[n-i-1][n-j-1]
			matrix[n-i-1][n-j-1] = matrix[j][n-i-1]
			matrix[j][n-i-1] = temp
		}
	}
}

// [i][j] => [j][col-i-1]
func rotateV1(matrix [][]int) {
	temp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		temp[i] = make([]int, len(matrix[0]))
	}

	for i, row := range matrix {
		for j, v := range row {
			temp[j][len(matrix)-i-1] = v
		}
	}

	copy(matrix, temp)
}

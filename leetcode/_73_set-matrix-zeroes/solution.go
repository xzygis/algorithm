package _73_set_matrix_zeroes

func setZeroes(matrix [][]int) {
	row := make([]bool, len(matrix))
	col := make([]bool, len(matrix[0]))
	for i, nums := range matrix {
		for j, num := range nums {
			if num == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}

	for i, nums := range matrix {
		for j := range nums {
			if row[i] || col[j] {
				matrix[i][j] = 0
			}
		}
	}
}

func setZeroesV0(matrix [][]int) {
	row := make([]int, len(matrix))
	for i := 0; i < len(row); i++ {
		row[i] = 1
	}

	col := make([]int, len(matrix[0]))
	for i := 0; i < len(col); i++ {
		col[i] = 1
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				row[i] = 0
				col[j] = 0
			}
		}
	}

	for i := 0; i < len(row); i++ {
		if row[i] == 0 {
			for j := 0; j < len(col); j++ {
				matrix[i][j] = 0
			}
		}
	}

	for i := 0; i < len(col); i++ {
		if col[i] == 0 {
			for j := 0; j < len(row); j++ {
				matrix[j][i] = 0
			}
		}
	}
}

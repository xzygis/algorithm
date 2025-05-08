package _47_search_a_2d_matrix

func searchMatrix(matrix [][]int, target int) bool {
	row, col := len(matrix), len(matrix[0])
	left, right := 0, row*col-1
	for left <= right {
		mid := (left + right) / 2
		i := mid / col
		j := mid % col
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return false
}

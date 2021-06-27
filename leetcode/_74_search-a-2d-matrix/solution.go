package _74_search_a_2d_matrix

func searchMatrix(matrix [][]int, target int) bool {
	rows := len(matrix)
	cols := len(matrix[0])
	left := 0
	right := rows * cols - 1
	for left <= right {
		mid := left + (right - left) / 2
		row := mid / cols
		col := mid % cols
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

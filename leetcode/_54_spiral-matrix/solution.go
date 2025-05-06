package _54_spiral_matrix

var directions = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func spiralOrder(matrix [][]int) []int {
	visited := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		visited[i] = make([]int, len(matrix[0]))
	}

	var ans []int
	i, j := 0, 0
	direction := 0
	for len(ans) < len(matrix)*len(matrix[0]) {
		if visited[i][j] == 0 {
			ans = append(ans, matrix[i][j])
			visited[i][j] = 1
		}

		nextI := i + directions[direction][0]
		nextJ := j + directions[direction][1]
		if nextI < 0 || nextI >= len(matrix) || nextJ < 0 || nextJ >= len(matrix[0]) || visited[nextI][nextJ] == 1 {
			direction = (direction + 1) % len(directions)
		}

		i += directions[direction][0]
		j += directions[direction][1]
	}

	return ans
}

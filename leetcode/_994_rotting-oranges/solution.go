package _994_rotting_oranges

var directions = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func orangesRotting(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	var queue []int
	count := 0 //1的数量
	depth := make(map[int]int)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 2 {
				code := i*cols + j //注意是i*cols+j，而不是i*j+j
				queue = append(queue, code)
				depth[code] = 0
			} else if grid[i][j] == 1 {
				count++
			}
		}
	}

	ans := 0
	for len(queue) > 0 {
		code := queue[0]
		queue = queue[1:]
		i, j := code/cols, code%cols
		for d := 0; d < len(directions); d++ {
			nextI := i + directions[d][0]
			nextJ := j + directions[d][1]
			nextCode := nextI*cols + nextJ //此处同上
			if nextI < 0 || nextI >= rows || nextJ < 0 || nextJ >= cols {
				continue
			}

			if grid[nextI][nextJ] != 1 {
				continue
			}

			count--
			grid[nextI][nextJ] = 2
			depth[nextCode] = depth[code] + 1
			queue = append(queue, nextCode)
			ans = depth[nextCode]
		}
	}

	// 由于只能往上下左右四个方向扩散，有可能有些1永远无法变成2
	if count > 0 {
		return -1
	}

	return ans
}

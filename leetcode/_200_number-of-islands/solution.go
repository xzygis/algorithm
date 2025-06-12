package _200_number_of_islands

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	rows := len(grid)
	cols := len(grid[0])

	numsOfDfs := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				numsOfDfs++
				dfs(grid, i, j)
			}
		}
	}

	return numsOfDfs
}

func dfs(grid [][]byte, i int, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return
	}

	if grid[i][j] == '0' {
		return
	}

	grid[i][j] = '0'
	dfs(grid, i-1, j)
	dfs(grid, i+1, j)
	dfs(grid, i, j-1)
	dfs(grid, i, j+1)
}

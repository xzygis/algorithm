package _79_word_search

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(board[0]))
	}

	for i, row := range board {
		for j := range row {
			if backtrack(board, i, j, word, 0, visited) {
				return true
			}
		}
	}

	return false
}

func backtrack(board [][]byte, i int, j int, word string, start int, visited [][]bool) bool {
	if start == len(word) {
		return true
	}

	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return false
	}

	if visited[i][j] {
		return false
	}

	if board[i][j] != word[start] {
		return false
	}

	visited[i][j] = true
	ans := backtrack(board, i+1, j, word, start+1, visited) ||
		backtrack(board, i, j+1, word, start+1, visited) ||
		backtrack(board, i-1, j, word, start+1, visited) ||
		backtrack(board, i, j-1, word, start+1, visited)
	visited[i][j] = false
	return ans
}

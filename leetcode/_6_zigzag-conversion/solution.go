package zigzagconversion

import "strings"

func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	rows := make([]string, numRows)
	row := 0
	forward := -1
	for _, ch := range s {
		rows[row] = rows[row] + string(ch)
		if row == 0 || row == numRows-1 {
			forward = -forward
		}

		row += forward
	}
	return strings.Join(rows, "")
}

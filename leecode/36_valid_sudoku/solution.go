package _6_valid_sudoku

/**
请你判断一个 9x9 的数独是否有效。只需要 根据以下规则 ，验证已经填入的数字是否有效即可。

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）
数独部分空格内已填入了数字，空白格用 '.' 表示。

注意：

一个有效的数独（部分已被填充）不一定是可解的。
只需要根据以上规则，验证已经填入的数字是否有效即可。
*/

func isValidSudoku(board [][]byte) bool {
	m := len(board)
	n := len(board[0])
	var cols [9][9]bool
	var rows [9][9]bool
	var blocks [9][9]bool
	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {
			if board[x][y] != '.' {
				num := board[x][y] - '1'
				blockIdx := x/3*3 + y/3
				if rows[x][num] || cols[y][num] || blocks[blockIdx][num] {
					return false
				}
				rows[x][num] = true
				cols[y][num] = true
				blocks[blockIdx][num] = true
			}
		}
	}
	return true
}

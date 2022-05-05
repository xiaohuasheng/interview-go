package backtrackingmethod

func solveSudoku(board [][]byte) {
	fill(0, 0, board)
}

func fill(i, j int, board [][]byte) bool {
	if j == 9 {
		i++
		j = 0
		if i == 9 {
			return true
		}
	}
	if board[i][j] != '.' {
		return fill(i, j+1, board)
	}

	for num := byte('1'); num <= byte('9'); num++ {
		if hasConflit(i, j, num, board) {
			continue
		}
		board[i][j] = num
		if fill(i, j+1, board) {
			return true
		}
		board[i][j] = '.'
	}
	return false
}

func hasConflit(r, c int, val byte, board [][]byte) bool {
	for i := 0; i < 9; i++ {
		if board[i][c] == val || board[r][i] == val {
			return true
		}
	}
	subRowStart := (r / 3) * 3
	subColStart := (c / 3) * 3

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[subRowStart+i][subColStart+j] == val {
				return true
			}
		}
	}
	return false
}

// 作者：xiao_ben_zhu
// 链接：https://leetcode-cn.com/problems/sudoku-solver/solution/shou-hua-tu-jie-jie-shu-du-hui-su-suan-fa-sudoku-s/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

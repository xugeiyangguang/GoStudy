package unitTestExcercise

func exist1(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	record := [][]int{}
	for i := 0; i < len(board); i++ {
		tmp := []int{}
		for j := 0; j < len(board[0]); j++ {
			tmp = append(tmp, 0)
		}
		record = append(record, tmp)
	}

	//record:=[len(board)][len(board[0])]int{}
	re := false
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				tmp := helper79(board, i, j, word, record)
				if tmp == true {
					return true
				}
			}
		}
	}
	return re
}

func helper79(board [][]byte, row, col int, word string, record [][]int) bool {
	if len(word) == 0 {
		return true
	}
	if row < 0 || col < 0 || row >= len(board) || col >= len(board[0]) || record[row][col] == 1 || board[row][col] != word[0] {
		return false
	}
	record[row][col] = 1
	re := helper79(board, row+1, col, word[1:], record) ||
		helper79(board, row-1, col, word[1:], record) ||
		helper79(board, row, col+1, word[1:], record) ||
		helper79(board, row, col-1, word[1:], record)
	if re == false {
		record[row][col] = 0
		return false
	}
	return true
}

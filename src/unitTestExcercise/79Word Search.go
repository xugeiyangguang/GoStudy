package unitTestExcercise

import "fmt"

func exist(board [][]byte, word string) bool {
	if board == nil {
		return false
	}
	if len(word) == 0 {
		return true
	}

	foot := [][]bool{}
	for i := 0; i < len(board); i++ {
		tmp := []bool{}
		for j := 0; j < len(board[0]); j++ {
			tmp = append(tmp, false)
		}
		foot = append(foot, tmp)
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			re := helperExist(board, i, j, foot, word, 0)
			if re == true {
				return true
			}
		}
	}
	return false
}

func helperExist(board [][]byte, row, col int, foot [][]bool, word string, cur int) bool {
	if cur == len(word) {
		return true
	}
	var re bool
	if foot[row][col] == true || board[row][col] != word[cur] {
		return false
	} else {
		foot[row][col] = true
		cur++
		if cur == len(word) {
			return true
		}
		if row-1 >= 0 {
			re = helperExist(board, row-1, col, foot, word, cur)
		}
		if row+1 < len(board) {
			re = re || helperExist(board, row+1, col, foot, word, cur)
		}
		if col-1 >= 0 {
			re = re || helperExist(board, row, col-1, foot, word, cur)
		}
		if col+1 < len(board[0]) {
			re = re || helperExist(board, row, col+1, foot, word, cur)
		}
		foot[row][col] = false
	}
	return re
}

func main() {
	a := [][]byte{{'C', 'A', 'A'}, {'A', 'A', 'A'}, {'B', 'C', 'D'}}

	fmt.Println(exist(a, "AAB"))
}

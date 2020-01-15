package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	//map嵌套需要对map进行初始化
	var row [9]map[byte]int
	for i := range row {
		row[i] = make(map[byte]int)
	}
	var col [9]map[byte]int
	for i := range col {
		col[i] = make(map[byte]int)
	}
	var sub [9]map[byte]int
	for i := range sub {
		sub[i] = make(map[byte]int)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			if board[i][j] < '0' || board[i][j] > '9' {
				return false
			}

			col[j][board[i][j]]++
			if col[j][board[i][j]] > 1 {
				return false
			}
			row[i][board[i][j]]++
			if row[i][board[i][j]] > 1 {
				return false
			}
			if j <= 2 && i <= 2 {
				sub[0][board[i][j]]++
				if sub[0][board[i][j]] > 1 {
					return false
				}
			} else if i >= 3 && i <= 5 && j <= 2 {
				sub[3][board[i][j]]++
				if sub[3][board[i][j]] > 1 {
					return false
				}
			} else if i >= 6 && j <= 2 {
				sub[6][board[i][j]]++
				if sub[6][board[i][j]] > 1 {
					return false
				}
			} else if i <= 2 && j >= 3 && j <= 5 {
				sub[1][board[i][j]]++
				if sub[1][board[i][j]] > 1 {
					return false
				}
			} else if i >= 3 && i <= 5 && j >= 3 && j <= 5 {
				sub[4][board[i][j]]++
				if sub[4][board[i][j]] > 1 {
					return false
				}
			} else if i >= 6 && j >= 3 && j <= 5 {
				sub[7][board[i][j]]++
				if sub[7][board[i][j]] > 1 {
					return false
				}
			} else if i <= 2 && j >= 6 {
				sub[2][board[i][j]]++
				if sub[2][board[i][j]] > 1 {
					return false
				}
			} else if i >= 3 && i <= 5 && j >= 6 {
				sub[5][board[i][j]]++
				if sub[5][board[i][j]] > 1 {
					return false
				}
			} else {
				sub[8][board[i][j]]++
				if sub[8][board[i][j]] > 1 {
					return false
				}
			}

		}
	}
	return true
}

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	fmt.Println(isValidSudoku(board))
}

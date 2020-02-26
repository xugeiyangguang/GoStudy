package unitTestExcercise

func gameOfLife(board [][]int) {
	reverse := make([][]int, 0)
	for i := range board {
		for j := range board[i] {
			one := 0
			if i-1 >= 0 && board[i-1][j] == 1 {
				one++
			}
			if i+1 < len(board) && board[i+1][j] == 1 {
				one++
			}
			if j-1 >= 0 && board[i][j-1] == 1 {
				one++
			}
			if j+1 < len(board[0]) && board[i][j+1] == 1 {
				one++
			}
			if i-1 >= 0 && j-1 >= 0 && board[i-1][j-1] == 1 {
				one++
			}
			if i-1 >= 0 && j+1 < len(board[0]) && board[i-1][j+1] == 1 {
				one++
			}
			if i+1 < len(board) && j-1 >= 0 && board[i+1][j-1] == 1 {
				one++
			}
			if i+1 < len(board) && j+1 < len(board[0]) && board[i+1][j+1] == 1 {
				one++
			}
			if board[i][j] == 1 {
				if one < 2 || one > 3 {
					reverse = append(reverse, []int{i, j})
				}
			} else {
				if one == 3 {
					reverse = append(reverse, []int{i, j})
				}
			}
		}
	}

	for i := range reverse {
		a := reverse[i][0]
		b := reverse[i][1]
		if board[a][b] == 1 {
			board[a][b] = 0
		} else {
			board[a][b] = 1
		}
	}

}

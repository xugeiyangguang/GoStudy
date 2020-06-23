package unitTestExcercise

func setZeroes(matrix [][]int) {
	row := make([]int, 0)
	col := make([]int, 0)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				row = append(row, i)
				col = append(col, j)
			}
		}
	}
	for j := 0; j < len(row); j++ {
		for i := 0; i < len(matrix[0]); i++ {
			matrix[row[j]][i] = 0
		}
	}
	for i := 0; i < len(col); i++ {
		for j := 0; j < len(matrix); j++ {
			matrix[j][col[i]] = 0
		}
	}

}

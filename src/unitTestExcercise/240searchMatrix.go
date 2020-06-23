package unitTestExcercise

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	return helper240(matrix, 0, len(matrix[0])-1, target)
}
func helper240(matrix [][]int, row, col, target int) bool {
	if row >= len(matrix) || col < 0 {
		return false
	}
	if matrix[row][col] == target {
		return true
	} else if matrix[row][col] < target {
		return helper240(matrix, row+1, col, target)
	} else {
		return helper240(matrix, row, col-1, target)
	}
}

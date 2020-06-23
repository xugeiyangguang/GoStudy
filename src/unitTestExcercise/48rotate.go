package unitTestExcercise

func rotate(matrix [][]int) {
	helper48(matrix, 0)
}
func helper48(matrix [][]int, start int) {
	if start > len(matrix)/2 {
		return
	}
	for i := 0; i < len(matrix); i++ {
		matrix[i][len(matrix)-1], matrix[len(matrix)-1][len(matrix)-1-i], matrix[len(matrix)-1-i][0], matrix[0][i] =
			matrix[0][i], matrix[i][len(matrix)-1], matrix[len(matrix)-1][len(matrix)-1-i], matrix[len(matrix)-1-i][0]
	}
	helper48(matrix, start+1)
}

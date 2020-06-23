package unitTestExcercise

func kthSmallest1(matrix [][]int, k int) int {
	row, col := len(matrix), len(matrix[0])
	left, right := matrix[0][0], matrix[row-1][col-1]
	for left < right {
		mid := (left + right) / 2
		tmp := helper378(matrix, mid)
		if tmp < k {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return right
}
func helper378(matrix [][]int, num int) int {
	re := 0
	for i := 0; i < len(matrix); i++ {
		j := 0
		for ; j < len(matrix[0]); j++ {
			if matrix[i][j] > num {
				break
			}
		}
		re += j
	}
	return re
}

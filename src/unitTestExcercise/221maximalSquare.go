package unitTestExcercise

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	index := 0
	dp := make([][]int, 0)
	for i := 0; i < len(matrix); i++ {
		tmp := []int{}
		for j := 0; j < len(matrix[0]); j++ {
			if i == 0 || j == 0 {
				if matrix[i][j] == '0' {
					tmp = append(tmp, 0)
				} else {
					tmp = append(tmp, 1)
					index = 1

				}

			} else {
				tmp = append(tmp, 0)
			}

		}
		dp = append(dp, tmp)

	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == '0' {
				dp[i][j] = 0
			} else {
				dp[i][j] = minThree(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
				index = max(index, dp[i][j])
			}
		}
	}

	return index * index

}

func minThree(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}

package unitTestExcercise

func minPathSum(grid [][]int) int {
	dp := make([]int, 0)
	dp = append(dp, grid[0][0])
	for i := 1; i < len(grid[0]); i++ {
		dp = append(dp, dp[i-1]+grid[0][i])
	}
	for i := 1; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if j == 0 {
				dp[j] = dp[j] + grid[i][j]
			} else {
				dp[j] = min(dp[j-1], dp[j]) + grid[i][j]
			}
		}
	}
	return dp[len(dp)-1]
}

package unitTestExcercise

func uniquePaths(m int, n int) int {
	dp := []int{}
	for i := 0; i < n; i++ {
		dp = append(dp, 1)
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] += dp[j-1]
		}
	}
	return dp[n-1]
}

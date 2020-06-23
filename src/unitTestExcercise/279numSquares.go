package unitTestExcercise

func numSquares(n int) int {
	dp := make([]int, n+1)
	cnt := 0
	for i := 1; i < n; i++ {
		dp[i] = i
		for j := 1; i > j*j; j++ {
			cnt = dp[i-j*j] + 1
			if cnt < dp[i] {
				dp[i] = cnt
			}
		}
	}
	return dp[n]
}

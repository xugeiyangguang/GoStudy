package unitTestExcercise

func maxProfit1(prices []int) int {
	n := len(prices)
	if n == 0 || n == 1 {
		return 0
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dp := [3]int{0, -prices[0], 0} //不持有，持有，冷却
	for i := 1; i < len(prices); i++ {
		dp[0], dp[1], dp[2] = max(dp[0], dp[1]+prices[i]), max(dp[1], dp[2]-prices[1]), dp[0]
	}
	return dp[0]
}

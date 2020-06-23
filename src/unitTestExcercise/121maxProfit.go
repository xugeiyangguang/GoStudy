package unitTestExcercise

func maxProfit2(prices []int) int {
	maxProfit := 0
	left, right := 0, 0

	for right < len(prices) {
		if prices[right] < prices[left] {
			left = right
		} else if cur := prices[right] - prices[left]; cur > maxProfit {
			maxProfit = cur
		}
		right++
	}
	return maxProfit
}

func max(a, b int) int {
	if a < b {
		return b

	} else {
		return a
	}
}

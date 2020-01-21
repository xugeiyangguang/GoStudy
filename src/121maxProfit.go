package main

import "fmt"

func maxProfit(prices []int) int {
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

func main() {
	price := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(price))

}

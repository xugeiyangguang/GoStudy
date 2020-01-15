package main

import "fmt"

func maxProfit(prices []int) int {
	var max int
	maxProfitCore(prices, 0, 0, &max)
	return max

}

func maxProfitCore(prices []int, start int, now int, max *int) {
	if start > len(prices)-2 {
		if now > *max {
			*max = now
		}
	}
	for i := start; i < len(prices)-1; i++ {
		for j := i + 1; j <= len(prices)-1; j++ {
			if prices[j] > prices[i] {
				maxProfitCore(prices, j+2, now+prices[j]-prices[i], max)
			}
		}
	}
	if now > *max {
		*max = now
	}
}

func main() {
	a := []int{1, 11, 2, 7, 4}
	fmt.Println(maxProfit(a))
}

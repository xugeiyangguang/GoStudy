package unitTestExcercise

import "math"

func coinChange1(coins []int, amount int) int {
	if amount < 1 {
		return 0
	}
	return helper322(coins, amount, make([]int, amount))
}

func helper322(coins []int, amount int, dp []int) int {
	if amount < 0 {
		return 0
	}
	if amount == 0 {
		return 0
	}
	if dp[amount-1] != 0 {
		return dp[amount-1]
	}
	minLen := math.MaxInt64
	for i := 0; i < len(coins); i++ {
		tmp := helper322(coins, amount-coins[i], dp)
		if tmp >= 0 && tmp < minLen {
			minLen = tmp + 1
		}
	}
	if minLen == math.MaxInt64 {
		dp[amount-1] = -1
	} else {
		dp[amount-1] = minLen
	}
	return dp[amount-1]
}

func coinChange(coins []int, amount int) int {
	if amount < 1 {
		return 0
	}
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}

		}
	}
	if dp[amount] > amount {
		return 0
	}
	return dp[amount]
}

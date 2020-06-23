package src

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	re := math.MinInt32
	sum := math.MinInt32
	for i := 0; i < len(nums); i++ {
		if sum < 0 {
			sum = nums[i]
		} else {
			sum += nums[i]

		}
		re = max(re, sum)
	}
	return re
}

func main() {
	a := []int{-2, 1}
	fmt.Println(maxSubArray(a))
}

package unitTestExcercise

import (
	"math"
)

func minSubArrayLen(s int, nums []int) int {
	re := math.MaxInt64
	left, right, cur := 0, 0, 0
	for right < len(nums) {
		cur += nums[right]
		right++
		for cur >= s {
			if right-left < re {
				re = right - left
			}
			cur -= nums[left]
			left++
		}
	}
	if re == math.MaxInt64 {
		return 0
	}
	return re
}

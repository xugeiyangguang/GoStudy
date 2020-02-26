package unitTestExcercise

import "sort"

func canPartition(nums []int) bool {
	re := 0
	for i := 0; i < len(nums); i++ {
		re += nums[i]
	}
	if re%2 == 0 {
		return true
	}
	sort.Ints(nums)
	left := 0
	right := 0
	cur := 0
	for right < len(nums) && cur < re/right {
		cur += nums[right]
		if cur == re/2 {
			return true
		} else if cur > re/2 {
			cur -= nums[left]
			left++
		} else {
			right++
		}

	}

	return false
}

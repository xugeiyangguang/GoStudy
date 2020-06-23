package unitTestExcercise

import "fmt"

func findTargetSumWays(nums []int, S int) int {
	if len(nums) == 1 {
		if S == nums[0] && S == -nums[0] {
			return 2
		} else if S == nums[0] || S == -nums[0] {
			return 1
		} else {
			return 0
		}
	}
	return findTargetSumWays(nums[1:], S-nums[0]) + findTargetSumWays(nums[1:], S+nums[0])
}

func helper494(nums []int, S int, re []byte) int {
	if len(nums) == 1 && S == nums[0] {
		fmt.Println(re)
		return 1
	}
	if len(nums) == 1 {
		return 0
	}
	add := helper494(nums[1:], S-nums[0], append(re, '+'))
	minus := helper494(nums[1:], S+nums[0], append(re, '-'))
	return add + minus
}

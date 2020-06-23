package unitTestExcercise

import "sort"

func nextPermutation(nums []int) {
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			for j := i; j <= len(nums); j++ {
				if j == len(nums) || nums[j] <= nums[i-1] && nums[j-1] > nums[i-1] {
					nums[i-1], nums[j-1] = nums[j-1], nums[i-1]
					sort.Ints(nums[i:])
					return
				}
			}
		}
	}
	sort.Ints(nums)
}

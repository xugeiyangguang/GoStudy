package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	re := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		start := i + 1
		end := len(nums) - 1
		for start < end {
			if nums[start]+nums[end] < 0-nums[i] {
				start++
			} else if nums[start]+nums[end] > 0-nums[i] {
				end--
			} else {
				tmp := []int{nums[i], nums[start], nums[end]}

				re = append(re, tmp)

				start++
				end--
				for start < end && nums[start-1] == nums[start] {
					start++
				}
				for start < end && nums[end+1] == nums[end] {
					end--
				}
			}
		}
	}
	return re
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	re := threeSum(nums)
	println(re)

}

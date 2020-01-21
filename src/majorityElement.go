package main

import "fmt"

func majorityElement(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	left := 0
	right := len(nums) - 1
	pos := sort(nums, left, right)
	for pos != len(nums)/2 {
		if pos < len(nums)/2 {
			pos = sort(nums, pos+1, right)
		} else {
			pos = sort(nums, left, pos-1)
		}
	}
	return nums[pos]

}

func sort(nums []int, left, right int) int {
	key := nums[left]
	for left < right {
		for left < right && nums[right] >= key {
			right--
		}
		nums[left] = nums[right]
		for left < right && nums[left] <= key {
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = key
	return left
}
func main() {
	nums := []int{3, 2, 3}
	fmt.Println(majorityElement(nums))
}

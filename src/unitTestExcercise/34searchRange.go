package unitTestExcercise

func searchRange(nums []int, target int) []int {
	start := findStart(nums, target)
	last := findLast(nums, target)
	return []int{start, last}
}

func findStart(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
		if mid-1 >= 0 && nums[mid] == target && nums[mid-1] != target ||
			mid-1 < 0 && nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func findLast(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
		if mid+1 < len(nums) && nums[mid] == target && nums[mid+1] != target ||
			mid+1 >= len(nums) && nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

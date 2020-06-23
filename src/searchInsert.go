package src

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	var mid int
	for left < right {
		mid = (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if nums[left] >= target {
		return left
	}
	return left + 1
}

func main() {
	nums := []int{1}
	println(searchInsert(nums, 1))
}

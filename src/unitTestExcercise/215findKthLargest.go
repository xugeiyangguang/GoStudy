package unitTestExcercise

func findKthLargest(nums []int, k int) int {
	cur := -1
	start, end := 0, len(nums)-1
	k = len(nums) - k
	for cur != k {
		cur = partition215(nums, start, end)
		if cur > k {
			end = cur - 1
		} else if cur < k {
			start = cur + 1
		} else {
			return nums[cur]
		}

	}
	return nums[cur]
}

func partition215(nums []int, start int, end int) int {
	left, right := start, end
	key := nums[start]
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

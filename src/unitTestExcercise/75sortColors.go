package unitTestExcercise

func sortColors(nums []int) {
	left := 0
	right := len(nums) - 1
	for left < len(nums)-1 && nums[left] == 0 {
		left++
	}
	for right >= 0 && nums[right] == 2 {
		right--
	}
	cur := left
	for cur <= right {
		if nums[cur] == 2 {
			nums[cur], nums[right] = nums[right], nums[cur]
			right--
		} else if nums[cur] == 0 {
			if cur == left {
				cur++
				left++
			} else {
				nums[left], nums[cur] = nums[cur], nums[left]
				cur++
				left++
			}
		} else {
			cur++
		}
	}
}

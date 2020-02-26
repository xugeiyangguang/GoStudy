package unitTestExcercise

func findDuplicate(nums []int) int {
	for i := 0; i < len(nums); {
		if nums[i] == i+1 {
			i++
		} else {
			if nums[nums[i]-1] == nums[i] {
				return nums[i]
			}
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	return 0
}

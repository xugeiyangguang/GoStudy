package unitTestExcercise

func canJump(nums []int) bool {
	var maxindex int
	for i := range nums {
		if i > maxindex {
			return false
		}
		maxindex = max(maxindex, i+nums[i])
	}
	return maxindex >= len(nums)-1
}

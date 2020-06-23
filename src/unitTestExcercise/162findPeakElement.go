package unitTestExcercise

func findPeakElement(nums []int) int {
	if len(nums) <= 1 {
		return len(nums) - 1
	}
	if nums[1] < nums[0] {
		return 0
	}
	if nums[len(nums)-2] < nums[len(nums)-1] {
		return len(nums) - 1
	}
	for i := 1; i < len(nums)-1; i++ {
		if nums[i-1] < nums[i] && nums[i+1] < nums[i] {
			return i
		}
	}

	return -1
}

//func findPeakElement(nums []int) int {
//	return findCore(nums, 0, len(nums)-1)
//}
//
//func findCore(nums []int, left, right int) int {
//	mid := (left + right) / 2
//	if mid == 0 && nums[mid+1] < nums[mid] || mid == len(nums)-1 && nums[mid-1] < nums[mid] ||
//		nums[mid-1] < nums[mid] && nums[mid+1] < nums[mid] {
//		return mid
//	}
//	r1 := findCore(nums, mid+1, right)
//	if r1 != -1 {
//		return r1
//	}
//	return findCore(nums, left, mid-1)
//}

package unitTestExcercise

func permute(nums []int) [][]int {
	re := [][]int{}
	helper46(nums, 0, &re)
	return re
}
func helper46(nums []int, start int, re *[][]int) {
	if start == len(nums)-1 {
		*re = append(*re, append([]int{}, nums...))
	}
	for i := start; i < len(nums); i++ {
		nums[i], nums[start] = nums[start], nums[i]
		helper46(nums, start+1, re)
		nums[i], nums[start] = nums[start], nums[i]
	}
}

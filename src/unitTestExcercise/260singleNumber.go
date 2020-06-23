package unitTestExcercise

func singleNumber(nums []int) []int {
	total := 0
	for i := range nums {
		total = total ^ nums[i]
	}
	pos := 1
	for pos&total == 0 {
		pos <<= 1
	}
	a := 0
	b := 0
	for i := range nums {
		if nums[i]&pos == 0 {
			a = a ^ nums[i]
		} else {
			b = b ^ nums[i]
		}
	}
	return []int{a, b}
}

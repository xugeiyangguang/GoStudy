package main

func moveZeroes(nums []int) {
	last := 0
	cur := 0
	for ; last < len(nums); cur++ {
		if cur >= len(nums) {
			nums[last] = 0
			last++
		} else if nums[cur] != 0 {
			nums[last] = nums[cur]
			last++
		}

	}
	/*
		for _, v := range nums {
			if v == 0 {
				continue
			}
			nums[last] = v
			last++
		}
		for last < len(nums) {
			nums[last] = 0
			last++
		}

	*/
}
func main() {
	re := []int{0, 1, 0, 3, 1, 2}
	moveZeroes(re)
	println(re)
}

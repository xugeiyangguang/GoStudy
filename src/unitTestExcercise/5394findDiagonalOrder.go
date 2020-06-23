package unitTestExcercise

func findDiagonalOrder(nums [][]int) []int {
	lens := 0
	for i := 0; i < len(nums); i++ {
		lens = max(len(nums[i]), lens)
	}
	re := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		x, y := i, 0
		for x >= 0 {
			if y < len(nums[x]) {
				re = append(re, nums[x][y])
			}
			x--
			y++
		}
	}
	for i := 1; i < lens; i++ {
		x, y := len(nums)-1, i
		for y < lens && x >= 0 {
			if y < len(nums[x]) {
				re = append(re, nums[x][y])
			}
			x--
			y++
		}
	}

	return re
}

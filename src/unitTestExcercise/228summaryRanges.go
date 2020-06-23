package unitTestExcercise

import "strconv"

func summaryRanges(nums []int) []string {
	re := make([]string, 0)
	if len(nums) == 0 {
		return re
	}
	//cur := strconv.Itoa(nums[0])
	start := nums[0]
	for i := 1; i <= len(nums); i++ {
		if i < len(nums) && nums[i] == nums[i-1]+1 {
			//cur = cur + "->" + strconv.Itoa(nums[i])
			continue
		} else {
			tmp := ""
			if nums[i-1] != start {
				tmp = strconv.Itoa(start) + "->" + strconv.Itoa(nums[i-1])
			} else {
				tmp = strconv.Itoa(start)
			}
			re = append(re, tmp)
			if i < len(nums) {
				start = nums[i]
			}
		}
	}
	return re
}

package unitTestExcercise

func majorityElement(nums []int) []int {
	candidate1, count1 := 0, 0
	candidate2, count2 := 1, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == candidate1 {
			count1++
		} else if nums[i] == candidate2 {
			count2++
		} else if count1 == 0 {
			candidate1 = nums[i]
			count1 = 1
		} else if count2 == 0 {
			candidate2 = nums[i]
			count2 = 1
		} else {
			count2--
			count1--
		}
	}
	check1, check2 := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == candidate1 {
			check1++
		}
		if nums[i] == candidate2 {
			check2++
		}
	}
	re := make([]int, 0)
	if check1 > len(nums)/3 {
		re = append(re, candidate1)
	}
	if check2 > len(nums)/3 {
		re = append(re, candidate2)
	}
	return re
}

package unitTestExcercise

func maxCoins(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var re int
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			re = nums[0]*nums[1] + maxCoins(nums[1:])
		} else if i == len(nums)-1 {
			re = Max(nums[len(nums)-2]*nums[len(nums)-1]+maxCoins(nums[:len(nums)-1]), re)
		} else {
			tmp := append(nums[0:i], nums[i+1:]...)
			re = Max(nums[i-1]*nums[i]*nums[i+1]+maxCoins(tmp), re)
		}
	}
	return re
}

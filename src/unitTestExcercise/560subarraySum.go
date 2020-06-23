package unitTestExcercise

func subarraySum(nums []int, k int) int {
	cur, re := nums[0], 0
	i, j := 0, 0
	for i < len(nums) {
		if cur == k {
			re++
			cur -= nums[i]
			i++
			j++
			if j == len(nums) {
				return re
			}
			cur += nums[j]
		} else if cur < k {
			j++
			if j == len(nums) {
				return re
			}
			cur += nums[j]
		} else {
			cur -= nums[i]
			i++
		}
	}
	return re
}

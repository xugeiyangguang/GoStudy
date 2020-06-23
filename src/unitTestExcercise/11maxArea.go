package unitTestExcercise

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	re := 0
	for left < right {
		tmp := 0
		if height[left] < height[right] {
			tmp = (right - left) * height[left]
			left++
		} else {
			tmp = (right - left) * height[right]
			right--
		}
		re = max(re, tmp)
	}
	return re
}

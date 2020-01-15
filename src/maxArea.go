package main

import "fmt"

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	re := 0
	var tmp int
	for left < right {

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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	re := maxArea(height)
	fmt.Println(re)
}

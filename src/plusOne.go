package main

func main() {
	digits := []int{9, 9}
	re := plusOne(digits)
	println(re)
}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] += 1
			return digits
		} else {
			digits[i] = 0
		}
	}
	re := []int{1}
	re = append(re, digits...)
	return re
}

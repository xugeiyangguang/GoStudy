package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	left := 0
	right := 0
	lens := 0
	record := make(map[byte]int)
	for right < len(s) {
		if record[s[right]] == 1 {
			for s[left] != s[right] {
				delete(record, s[left])
				left++
			}
			left++
		} else {
			record[s[right]] = 1

		}
		right++
		lens = max(right-left, lens)
	}
	return lens
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	re := lengthOfLongestSubstring("aaaaa")
	fmt.Println(re)
}

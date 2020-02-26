package unitTestExcercise

func findAnagrams(s string, p string) []int {
	re := make([]int, 0)
	pMap := make(map[byte]int, 0)
	for i := 0; i < len(p); i++ {
		pMap[p[i]]++
	}
	left := 0
	right := 0
	for left <= len(s)-len(p) {
		for right < len(s) && pMap[s[right]] > 0 {
			pMap[s[right]]--
			if right-left == len(p)-1 {
				re = append(re, left)
				pMap[s[left]]++
				left++
			}
			right++
		}
		for left <= len(s)-len(p) && s[left] != s[right] {
			pMap[s[left]]++
			left++
		}
		left++
		right++
	}
	return re
}

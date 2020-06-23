package unitTestExcercise

func longestSubstring(s string, k int) int {
	if len(s) < k {
		return 0
	}
	record := make(map[byte]int)
	for i, _ := range s {
		record[s[i]]++
	}
	for i, _ := range s {
		if record[s[i]] < k {
			left := longestSubstring(s[:i], k)
			for i < len(s) && record[s[i]] < k {
				i++
			}
			right := longestSubstring(s[i:], k)
			return Max(left, right)
		}
	}
	return len(s)

}

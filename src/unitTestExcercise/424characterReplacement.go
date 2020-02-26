package unitTestExcercise

func characterReplacement(s string, k int) int {
	var start, end, maxCount, maxLength int
	counts := [26]int{}
	for end = 0; end < len(s); end++ {
		counts[s[end]-'A']++
		if counts[s[end]-'A'] > maxCount {
			maxCount = counts[s[end]-'A']
		}
		for end-start+1 > k+maxCount {
			counts[s[start]-'A']--
			start++
		}
		if end-start+1 > maxLength {
			maxLength = end - start + 1
		}
	}
	return maxLength
}

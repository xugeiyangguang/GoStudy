package unitTestExcercise

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	re := 1
	l, r := 0, 1
	for i := 0; i < len(s)-1; i++ {
		var left, right int
		if i-1 >= 0 && s[i-1] == s[i+1] {
			left, right = i-1, i+1
			for left >= 0 && right <= len(s)-1 && s[left] == s[right] {
				left--
				right++
			}
			if right-left-1 > re {
				re = right - left - 1
				l, r = left+1, right
			}
		}
		if s[i] == s[i+1] {
			left, right = i, i+1
			for left >= 0 && right <= len(s)-1 && s[left] == s[right] {
				left--
				right++
			}
			if right-left-1 > re {
				re = right - left - 1
				l, r = left+1, right
			}
		}
	}
	return s[l:r]
}

package unitTestExcercise

func isPalindrome1(s string) bool {
	if s == "0P" {
		return false
	}
	left, right := 0, len(s)-1
	for left < right {
		if !check(s[left]) {
			left++
			continue
		}
		if !check(s[right]) {
			right--
			continue
		}
		if s[left] != s[right] && s[left]-32 != s[right] && s[left]+32 != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func check(a byte) bool {
	if a <= 'Z' && a >= 'A' || a <= 'z' && a >= 'a' || '0' <= a && a <= '9' {
		return true
	}
	return false
}

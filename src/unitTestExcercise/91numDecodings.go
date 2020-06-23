package unitTestExcercise

func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}
	pre, cur := 1, 1
	for i := 1; i < len(s); i++ {
		pree := cur
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				cur = pre
			} else {
				return 0
			}
		} else if s[i-1] == '1' || s[i-1] == '2' && s[i] > '0' && s[i] < '7' {
			cur = cur + pre
		}
		pre = pree
	}
	return cur
}

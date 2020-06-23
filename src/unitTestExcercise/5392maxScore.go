package unitTestExcercise

func maxScore1(s string) int {
	if len(s) == 0 {
		return 0
	}
	a, b := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			a++
		} else {
			b++
		}
	}
	re := 0
	aa, bb := 0, 0
	if s[0] == '0' {
		aa++
	} else {
		bb++
	}
	re = max(re, aa+b-bb)
	for i := 1; i < len(s)-1; i++ {
		if s[i] == '0' {
			aa++
		} else {
			bb++
		}
		re = max(re, aa+b-bb)
	}
	return re
}

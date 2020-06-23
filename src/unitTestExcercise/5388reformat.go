package unitTestExcercise

func reformat(s string) string {
	number := make([]byte, 0)
	letter := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] <= '9' && s[i] >= 0 {
			number = append(number, s[i])
		} else {
			letter = append(letter, s[i])
		}
	}
	if len(number)-len(letter) > 1 || len(letter)-len(number) > 1 {
		return ""
	}
	re := ""

	if len(number) > len(letter) {
		re += string(number[0])
		for i := 0; i < len(letter); i++ {
			re += string(letter[i])
			re += string(number[i+1])
		}
	} else if len(number) < len(letter) {
		re += string(letter[0])
		for i := 0; i < len(number); i++ {
			re += string(number[i])
			re += string(letter[i+1])
		}
	} else {
		for i := 0; i < len(letter); i++ {
			re += string(letter[i])
			re += string(number[i])
		}
	}
	return re
}

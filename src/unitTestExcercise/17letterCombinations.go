package unitTestExcercise

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	phone := make(map[byte][]string)
	phone['2'] = []string{"a", "b", "c"}
	phone['3'] = []string{"d", "e", "f"}
	phone['4'] = []string{"g", "h", "i"}
	phone['5'] = []string{"j", "k", "l"}
	phone['6'] = []string{"m", "n", "o"}
	phone['7'] = []string{"p", "q", "r", "s"}
	phone['8'] = []string{"t", "u", "v"}
	phone['9'] = []string{"w", "x", "y", "z"}

	re := phone[digits[0]]

	for i := 1; i < len(digits); i++ {
		tmp := phone[digits[i]]
		new_re := make([]string, 0)
		for j := 0; j < len(re); j++ {
			for k := 0; k < len(tmp); k++ {
				new_re = append(new_re, re[j]+tmp[k])
			}
		}
		re = new_re
	}
	return re
}

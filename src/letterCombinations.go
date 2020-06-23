package src

import "fmt"

func letterCombinations(digits string) []string {
	re := []string{}
	if len(digits) == 0 {
		return re
	}
	phone := make(map[byte][]byte)
	phone['2'] = []byte{'a', 'b', 'c'}
	phone['3'] = []byte{'d', 'e', 'f'}
	phone['4'] = []byte{'g', 'h', 'i'}
	phone['5'] = []byte{'j', 'k', 'l'}
	phone['6'] = []byte{'m', 'n', 'o'}
	phone['7'] = []byte{'p', 'q', 'r', 's'}
	phone['8'] = []byte{'t', 'u', 'v'}
	phone['9'] = []byte{'w', 'x', 'y', 'z'}

	one := phone[digits[0]]
	for i := range one {
		re = append(re, string(one[i]))
	}

	for i := 1; i < len(digits); i++ {
		tmp_re := []string{}
		for j := 0; j < len(re); j++ {
			for k := 0; k < len(phone[digits[i]]); k++ {
				tmp := re[j] + string(phone[digits[i]][k])
				tmp_re = append(tmp_re, tmp)
			}
		}
		re = tmp_re
	}
	return re
}

func main() {
	a := ""
	re := letterCombinations(a)
	fmt.Println(re)

}

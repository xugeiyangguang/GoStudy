package unitTestExcercise

import (
	"strconv"
)

func decodeString(s string) string {
	if len(s) <= 1 {
		return s
	}

	stack := make([]string, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == ']' {
			str := ""
			for stack[len(stack)-1] != "[" {
				str = stack[len(stack)-1] + str
				stack = stack[:len(stack)-1]
			}
			times := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			count, _ := strconv.Atoi(times)
			tmp := ""
			for j := 0; j < count; j++ {
				tmp = tmp + str
			}
			stack = append(stack, tmp)
		} else if s[i] <= '9' && s[i] >= '0' {
			start := i
			for s[i] <= '9' && s[i] >= '0' {
				i++
			}
			stack = append(stack, s[start:i])
			i--
		} else if s[i] <= 'z' && s[i] >= 'a' || s[i] <= 'Z' && s[i] >= 'A' {
			start := i
			for i < len(s) && (s[i] <= 'z' && s[i] >= 'a' || s[i] <= 'Z' && s[i] >= 'A') {
				i++
			}
			stack = append(stack, s[start:i])
			i--
		} else {
			stack = append(stack, "[")
		}
	}
	re := ""
	for i := 0; i < len(stack); i++ {
		re += stack[i]
	}
	return re
}

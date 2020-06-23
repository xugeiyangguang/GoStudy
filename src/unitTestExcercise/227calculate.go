package unitTestExcercise

func calculate(s string) int {
	stack := make([]int, 0)
	re, cur := 0, 0
	op := '+'
	for i := range s {
		if s[i] >= '0' && s[i] <= '9' {
			cur = cur*10 + int(s[i]-'0')
		}
		if s[i] == '+' || s[i] == '-' || s[i] == '*' || s[i] == '/' || i == len(s)-1 {
			if op == '+' {
				stack = append(stack, cur)
			} else if op == '-' {
				stack = append(stack, -cur)
			} else if op == '*' {
				tmp := cur * stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, tmp)
			} else if op == '/' {
				tmp := stack[len(stack)-1] / cur
				stack = stack[:len(stack)-1]
				stack = append(stack, tmp)
			}
			op = rune(s[i])
			cur = 0
		}

	}
	for _, v := range stack {
		re += v
	}
	return re
}

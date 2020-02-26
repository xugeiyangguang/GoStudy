package unitTestExcercise

import (
	"strconv"
	"strings"
)

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for i := 0; i < len(tokens); i++ {
		tmp := tokens[i]
		if len(tmp) == 1 && strings.ContainsAny(tmp, "+-*/") {
			a := stack[len(stack)-1]
			b := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			if tmp == "+" {
				stack = append(stack, a+b)
			} else if tmp == "-" {
				stack = append(stack, b-a)
			} else if tmp == "*" {
				stack = append(stack, a*b)
			} else {
				stack = append(stack, b/a)
			}
		} else {
			num, _ := strconv.Atoi(tmp)
			stack = append(stack, num)
		}
	}
	return stack[0]
}

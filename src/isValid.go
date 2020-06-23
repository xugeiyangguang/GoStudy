package src

func isValid(s string) bool {
	stack := make([]rune, len(s))
	top := 0
	for _, v := range s {
		switch v {
		case '{':
			stack[top] = '}'
			top++
			break
		case '[':
			stack[top] = ']'
			top++
			break
		case '(':
			stack[top] = ')'
			top++
			break
		default:
			if top == 0 || stack[top-1] != v {
				return false
			}
			top--
			break
		}

	}
	return top == 0
}

func main() {
	s := "()"
	println(isValid(s))
}

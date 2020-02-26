package unitTestExcercise

func removeKdigits(num string, k int) string {
	if len(num) == k {
		return "0"
	}
	stack := make([]byte, 0)
	i := 0
	for ; i < len(num); i++ {
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > num[i] {
			k--
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num[i])
	}
	stack = stack[:len(stack)-k]
	for j := 0; j < len(stack); j++ {
		if stack[j] != '0' {
			return string(stack[j:])
		}
	}
	return "0"
}

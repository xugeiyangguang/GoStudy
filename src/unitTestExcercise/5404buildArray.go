package unitTestExcercise

func buildArray(target []int, n int) []string {
	re := []string{}
	cur := 0
	for i := 1; i <= n; i++ {
		if target[cur] == i {
			re = append(re, "push")
			cur++
		} else if target[cur] > i {
			re = append(re, "push")
			re = append(re, "pop")
		}
		if cur == len(target) {
			break
		}
	}
	return re
}

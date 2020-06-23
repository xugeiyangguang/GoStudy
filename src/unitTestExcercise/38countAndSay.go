package unitTestExcercise

import "strconv"

func countAndSay(n int) string {
	cur, next := "1", "1"
	for i := 1; i <= n; i++ {
		start, j := 0, 1
		cur = next
		next = ""
		for ; j <= len(cur); j++ {
			if j == len(cur) || cur[j] != cur[start] {
				num := j - start
				next += strconv.Itoa(num) + string(cur[start])
				start = j
			}
		}
	}
	return cur
}

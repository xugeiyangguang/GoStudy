package src

import "strconv"

func main() {
	println(countAndSay(5))

}
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	pre := "1"
	for i := 2; i <= n; i++ {
		pre = core(pre)
	}
	return pre
}

func core(pre string) string {
	start := 0
	end := 0
	re := ""
	for end <= len(pre) {
		for end < len(pre) && (pre)[start] == (pre)[end] {
			end++
		}
		re = re + strconv.Itoa(end-start) + string((pre)[start])
		start = end
		end++
	}
	return re
}

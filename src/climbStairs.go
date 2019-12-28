package main

func climbStairs(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 1, 1
	var re int
	for i := 2; i <= n; i++ {
		re = a + b
		a = b
		b = re
	}
	return re
}

func main() {
	println(climbStairs(3))
}

package src

import "fmt"

func generateParenthesis(n int) []string {
	var re []string
	generateParenthesisCore(n, 0, 0, "", &re)
	return re
}

func generateParenthesisCore(n int, left int, right int, cur string, re *[]string) {
	if n == left {
		for right != n {
			cur += ")"
			right++
		}
		*re = append(*re, cur)
		return
	}
	if left > right {
		generateParenthesisCore(n, left, right+1, cur+")", re)
	}
	generateParenthesisCore(n, left+1, right, cur+"(", re)
}

func main() {
	re := generateParenthesis(1)
	fmt.Println(re)
}

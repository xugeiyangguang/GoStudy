package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	origin := x
	var a int
	for origin > 0 {
		a = a*10 + origin%10
		origin /= 10
	}
	if x == a {
		return true
	}
	return false
}
func main() {
	fmt.Print(isPalindrome(12))

}

package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {

	a := strings.Fields(s)
	if len(a) == 0 {
		return ""
	}
	re := ""
	for i := len(a) - 1; i >= 0; i-- {

		re += a[i] + " "
	}
	return re[0 : len(re)-1]
}

func main() {
	fmt.Println(reverseWords("  hello world!  "))
}

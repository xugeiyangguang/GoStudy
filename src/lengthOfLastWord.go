package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	ss := strings.Fields(s)
	if len(ss) == 0 {
		return 0
	}
	return len(ss[len(ss)-1])
}

func main() {
	re := lengthOfLastWord(" ")
	fmt.Println(re)
}

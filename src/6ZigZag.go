package main

import "fmt"

func convert(s string, numRows int) string {
	strs := make([]string, numRows)
	for i := 0; i < len(s); {
		for j := 0; j < numRows && i < len(s); i++ {
			strs[j] += string(s[i])
			j++
		}
		for j := numRows - 2; j > 0 && i < len(s); i++ {
			strs[j] += string(s[i])
			j--
		}
	}
	re := ""
	for i := range strs {
		re += strs[i]
	}
	return re
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 4))
}

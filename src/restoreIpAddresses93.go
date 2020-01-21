package main

import (
	"fmt"
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	re := make([]string, 0)
	parts := make([]string, 4)
	for i := 1; i < min(4, len(s)); i++ {
		parts[0] = s[0:i]
		if !invalid(parts[0]) {
			continue
		}
		for j := 1; j < min(4, len(s)-i); j++ {
			parts[1] = s[i : i+j]
			if !invalid(parts[1]) {
				continue
			}
			for k := 1; k < min(4, len(s)-i-j); k++ {
				parts[2] = s[i+j : k+i+j]
				if !invalid(parts[2]) {
					continue
				}
				parts[3] = s[k+i+j:]
				if !invalid(parts[3]) {
					continue
				}
				re = append(re, strings.Join(parts, "."))
			}
		}
	}
	return re
}

func invalid(part string) bool {
	if len(part) != 1 && part[0] == '0' {
		return false
	}
	tmp, _ := strconv.Atoi(part)
	if tmp > 255 || tmp < 0 {
		return false
	}
	return true
}

//func min(a, b int) int {
//	if a > b {
//		return b
//	} else {
//		return a
//	}
//}

func main() {
	re := restoreIpAddresses("0000")
	fmt.Println(re)
}

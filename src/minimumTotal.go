package main

import (
	"fmt"
	"math"
)

func minimumTotal(triangle [][]int) int {
	re := math.MaxInt64
	minimCore(triangle, 0, 0, triangle[0][0], &re)
	return re

}

func minimCore(triangle [][]int, row int, col int, cur int, re *int) {
	if row == len(triangle)-1 {
		*re = min(*re, cur)
		return
	}

	minimCore(triangle, row+1, col, cur+triangle[row+1][col], re)
	minimCore(triangle, row+1, col+1, cur+triangle[row+1][col+1], re)
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	triangle := [][]int{{2}}
	fmt.Println(minimumTotal(triangle))
}

package main

import "fmt"

func combine(n int, k int) [][]int {
	re := make([][]int, 0)
	helper(1, n, k, []int{}, &re)
	return re
}

func helper(cur int, n int, k int, tmp []int, re *[][]int) {
	if len(tmp) == k {
		*re = append(*re, append([]int{}, tmp...))
		return
	}
	for i := cur; i <= n; i++ {
		helper(i+1, n, k, append(tmp, i), re)
	}
}

func main() {
	re := combine(4, 3)
	fmt.Println(re)
}

package unitTestExcercise

func combine(n int, k int) [][]int {
	re := make([][]int, 0)
	helper77(1, n, k, []int{}, &re)
	return re
}

func helper77(cur int, n int, k int, tmp []int, re *[][]int) {
	if len(tmp) == k {
		*re = append(*re, append([]int{}, tmp...))
		return
	}
	for i := cur; i <= n; i++ {
		helper77(i+1, n, k, append(tmp, i), re)
	}
}

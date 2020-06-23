package unitTestExcercise

func numTrees(n int) int {
	if n == 0 {
		return 1
	}
	memo := make(map[int]int)
	re := 0
	for i := 1; i <= n; i++ {
		if _, ok := memo[i-1]; !ok {
			memo[i-1] = numTrees(i - 1)
		}
		if _, ok1 := memo[n-i]; !ok1 {
			memo[n-i] = numTrees(n - i)
		}
		re += memo[i-1] * memo[n-i]
	}
	return re
}

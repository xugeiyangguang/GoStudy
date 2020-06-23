package src

func countPrimes(n int) int {
	noPrime := make([]int, n)
	re := 0
	for i := 2; i < n; i++ {
		if noPrime[i] == 0 {
			re++
			for j := i + i; j < n; j = j + i {
				noPrime[j] = 1
			}
		}
	}
	return re
}

func main() {
	println(countPrimes(10))
}

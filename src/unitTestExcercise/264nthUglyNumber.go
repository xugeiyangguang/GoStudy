package unitTestExcercise

import "math"

func nthUglyNumber(n int) int {
	min := math.MaxInt64
	a := []int{1}
	p2, p3, p5 := 0, 0, 0
	for i := 1; i < n; i++ {
		min = threeMin(a[p2]*2, a[p3]*3, a[p5]*5)
		if i == n-1 {
			return min
		}
		a = append(a, min)
		for a[p2]*2 <= min {
			p2++
		}
		for a[p3]*3 <= min {
			p3++
		}
		for a[p5]*5 <= min {
			p5++
		}
	}
	return a[len(a)-1]
}
func threeMin(args ...int) int {
	min := math.MaxInt32
	for _, v := range args {
		if v < min {
			min = v
		}
	}
	return min
}

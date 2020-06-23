package unitTestExcercise

func getSum(a int, b int) int {
	for b != 0 {
		sum := a ^ b
		carray := a & b
		carray <<= 1
		a = sum
		b = carray
	}
	return a
}

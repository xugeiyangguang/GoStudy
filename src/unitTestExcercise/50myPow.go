package unitTestExcercise

func myPow(x float64, n int) float64 {

	if n < 0 {
		n, x = -n, 1/x
	}
	re := 1.0
	for n > 0 {
		if n&1 == 1 {
			re *= x
		}
		x, n = x*x, n>>1
	}
	return re
}

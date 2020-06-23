package unitTestExcercise

import "strconv"

func fractionToDecimal(numerator int, denominator int) string {
	re := ""
	a := numerator / denominator
	b := numerator % denominator
	if b != 0 {
		re += strconv.Itoa(a) + "."
	}
	last := -1
	for b != 0 {
		a = b * 10 / denominator
		b = b * 10 % denominator

		if last != -1 && b == last {
			re += "(" + strconv.Itoa(last) + ")"
			return re
		}
		re += strconv.Itoa(a)
		last = b
	}
	return re
}

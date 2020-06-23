package unitTestExcercise

import "math"

func reverseBits(num uint32) uint32 {
	index := 0
	re := uint32(0)
	for num != 0 {
		tmp := num & 1
		if tmp == 0 {
			re = re * 2
		} else {
			re = re*2 + 1
		}
		index++
	}
	if index != 32 {
		re = uint32(math.Pow(float64(re), float64(32-index))) * re
	}
	return re
}

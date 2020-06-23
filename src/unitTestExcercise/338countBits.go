package unitTestExcercise

func countBits(num int) []int {
	re := []int{0}
	if num == 0 {
		return re
	}
	for i := 1; i <= num; i++ {
		tmp := len(re)
		for j := 0; j < tmp && i <= num; j++ {
			re = append(re, re[j]+1)
			i++
		}
		i--
	}
	return re
}

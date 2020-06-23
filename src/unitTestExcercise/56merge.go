package unitTestExcercise

func merge(intervals [][]int) [][]int {
	if intervals == nil || len(intervals) == 0 {
		return intervals
	}
	sortArray(intervals)
	re := make([][]int, 0)
	//	re = append(re, intervals[0])
	for i := 0; i < len(intervals); i++ {
		flag := 0
		for j := 0; j < len(re); j++ {
			if re[j][1] >= intervals[i][0] {
				re[j][1] = max(re[j][1], intervals[i][1])
				flag = 1
				break
			}
		}
		if flag == 0 {
			re = append(re, intervals[i])
		}
	}
	return re
}

func sortArray(a [][]int) {
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-1; j++ {
			if a[j][0] > a[j+1][0] {
				a[j][0], a[j+1][0] = a[j+1][0], a[j][0]
				a[j][1], a[j+1][1] = a[j+1][1], a[j][1]
			}
		}
	}
}

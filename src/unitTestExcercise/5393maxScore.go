package unitTestExcercise

func maxScore(cardPoints []int, k int) int {
	record := make([][]int, 0)
	for i := 0; i < len(cardPoints); i++ {
		tmp := []int{}
		for j := 0; j < len(cardPoints); j++ {
			tmp = append(tmp, 0)
		}
		record = append(record, tmp)
	}
	return helper5393(cardPoints, 0, len(cardPoints)-1, k, record)
}
func helper5393(cardPoints []int, start, end int, k int, record [][]int) int {
	if k == 1 {
		return max(cardPoints[start], cardPoints[end])
	}
	if k == 0 {
		return 0
	}
	if record[start][end] != 0 {
		return record[start][end]
	}
	re := 0
	if len(cardPoints) == k {
		for i := 0; i < len(cardPoints); i++ {
			re += cardPoints[i]
		}
		return re
	}

	left := helper5393(cardPoints, start+1, end, k-1, record) + cardPoints[start]
	right := helper5393(cardPoints, start, end-1, k-1, record) + cardPoints[end]
	record[start][end] = max(left, right)
	return record[start][end]
}

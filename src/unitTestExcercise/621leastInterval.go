package unitTestExcercise

func leastInterval(tasks []byte, n int) int {
	record := make(map[byte]int)
	max := 0
	count := 0
	for _, v := range tasks {
		record[v]++
		if record[v] > max {
			count = 1
			max = record[v]
		} else if record[v] == max {
			count++
		}
	}
	return Max((max-1)*(n+1)+count, len(tasks))

}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

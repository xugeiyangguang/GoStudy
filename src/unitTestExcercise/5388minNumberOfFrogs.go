package unitTestExcercise

func minNumberOfFrogs(croakOfFrogs string) int {
	if len(croakOfFrogs)%5 != 0 {
		return -1
	}
	record := [5]int{}
	re := 0
	for i := 0; i < len(croakOfFrogs); i++ {
		if croakOfFrogs[i] == 'c' {
			if record[4] > 0 {
				record[4]--
			} else {
				re++
			}
			record[0]++
		} else if croakOfFrogs[i] == 'r' {
			record[0]--
			record[1]++
		} else if croakOfFrogs[i] == 'o' {
			record[1]--
			record[2]++
		} else if croakOfFrogs[i] == 'a' {
			record[2]--
			record[3]++
		} else {
			record[3]--
			record[4]++
		}
		if record[0] < 0 || record[1] < 0 || record[2] < 0 || record[3] < 0 {
			return -1
		}
	}
	if record[0] != 0 || record[1] != 0 || record[2] != 0 || record[3] != 0 {
		return -1
	}
	return re
}

package unitTestExcercise

func lengthOfLongestSubstring(s string) int {
	record := make(map[rune]int)
	re := 1
	start, end := 0, 0
	for i, v := range s {
		if record[v] == 1 {
			j := start
			for int32(s[j]) != v {
				delete(record, int32(s[j]))
				j++

			}
			start = j + 1
		} else {
			record[v] = 1
		}
		end = i
		re = Max(re, end-start+1)
	}
	return re
}

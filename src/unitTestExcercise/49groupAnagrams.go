package unitTestExcercise

func groupAnagrams1(strs []string) [][]string {
	re := make([][]string, 0)
	record := make(map[int]map[byte]int)
	zero := 0
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) == 0 {
			zero++
			continue
		}
		curMap := make(map[byte]int)
		for j := range strs[i] {
			curMap[strs[i][j]]++
		}
		j := 0
		for ; j < len(record); j++ {
			tag := 0
			if len(curMap) == len(record[j]) {
				for k, v := range curMap {
					if v != record[j][k] {
						tag = 1
						break
					}
				}
				if tag == 0 {
					re[j] = append(re[j], strs[i])
					break
				}
			}
		}
		if j == len(record) {
			record[j] = curMap
			re = append(re, []string{strs[i]})
		}
	}
	tmp := []string{}
	for zero != 0 {
		tmp = append(tmp, "")
		zero--
	}
	if len(tmp) != 0 {
		re = append(re, tmp)
	}
	return re
}

func groupAnagrams(words []string) [][]string {
	cache := make(map[[26]byte]int)
	result := make([][]string, 0)
	for i := range words {
		list := [26]byte{}
		for j := range words[i] {
			list[words[i][j]-'a']++
		}
		if idx, ok := cache[list]; ok {
			result[idx] = append(result[idx], words[i])
		} else {
			result = append(result, []string{words[i]})
			cache[list] = len(result) - 1
		}
	}
	return result
}

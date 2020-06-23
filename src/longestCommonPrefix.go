package src

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 || len(strs[0]) == 0 {
		return ""
	}
	end := len(strs[0])
	var i int
	for _, v := range strs {
		i = 0
		for i < len(v) && i < end {
			if strs[0][i] != v[i] {
				break
			}
			i++
		}
		if i == 0 {
			return ""
		}
		if i < end {
			end = i
		}
	}

	return strs[0][0:end]
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	println(longestCommonPrefix(strs))
}

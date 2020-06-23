package src

import "fmt"

func isIsomorphic(s string, t string) bool {
	mapp := make(map[rune][]int, len(s))
	mappp := make(map[rune][]int, len(t))
	for i, v := range s {
		mapp[v] = append(mapp[v], i)
	}
	for i, v := range t {
		mappp[v] = append(mappp[v], i)
	}
	for i, _ := range t {
		tmp1 := mapp[rune(s[i])]
		tmp2 := mappp[rune(t[i])]
		if len(tmp2) != len(tmp1) {
			return false
		}
		for j := 0; j < len(tmp1); j++ {
			if tmp2[j] != tmp1[j] {
				return false
			}
		}
	}
	return true
}

func main() {
	s := "aba"
	t := "cdd"
	re := isIsomorphic(s, t)
	fmt.Print(re)
}

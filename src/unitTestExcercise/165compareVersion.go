package unitTestExcercise

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	s1 := strings.Split(version1, ".")
	s2 := strings.Split(version2, ".")
	a1 := make([]int, 0)
	a2 := make([]int, 0)
	for i := 0; i < len(s1); i++ {
		tmp, _ := strconv.Atoi(s1[i])
		a1 = append(a1, tmp)
	}
	for i := 0; i < len(s2); i++ {
		tmp, _ := strconv.Atoi(s2[i])
		a2 = append(a2, tmp)
	}
	i := 0
	for ; i < len(a1) && i < len(a2); i++ {
		if a1[i] > a2[i] {
			return 1
		} else if a1[i] < a2[i] {
			return -1
		}
	}
	for i < len(a1) {
		if a1[i] != 0 {
			return 1
		}
		i++
	}
	for i < len(a2) {
		if a2[i] != 0 {
			return -1
		}
		i++
	}
	return 0
}

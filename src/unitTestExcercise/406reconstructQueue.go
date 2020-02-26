package unitTestExcercise

import (
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] > people[j][1]
		}
		return people[i][0] < people[j][0]
	})
	re := make([][]int, len(people))
	for _, p := range people {
		blank := 0
		k := p[1]
		for i := range re {
			if re[i] == nil {
				if blank == k {
					re[i] = p
				}
				blank++
			}
		}
	}
	return re
}

/*func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool{
		if people[i][0] == people[j][0] {
			return people[i][1] > people[j][1]
		}

		return people[i][0] < people[j][0]
	})

	ret := make([][]int, len(people))
	for _, p := range people {
		blanks := 0
		k := p[1]
		for i := range ret {
			if ret[i] == nil {
				if blanks == k {
					ret[i] = p
					break
				}
				blanks++
			}
		}
	}

	return ret
}*/

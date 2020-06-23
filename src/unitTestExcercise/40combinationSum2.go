package unitTestExcercise

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	re := make([][]int, 0)
	sort.Ints(candidates)
	helper(candidates, []int{}, target, &re)
	return re
}
func helper(candidates []int, record []int, target int, re *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		*re = append(*re, append([]int{}, record...))
		return
	}
	for i := 0; i < len(candidates); i++ {
		if i-1 < 0 && candidates[i] <= target || i-1 >= 0 && candidates[i] != candidates[i-1] {
			helper(candidates[i+1:], append(record, candidates[i]), target-candidates[i], re)
		}
	}
}

package main

import (
	"fmt"
)

/*func combinationSum(candidates []int, target int) [][]int {
	var cur []int
	var re [][]int
	sort.Ints(candidates)
	combinationSumCore(candidates, 0, target, &cur, &re)
	return re
}

func combinationSumCore(candidates []int, start int, target int, cur *[]int, re *[][]int) {

	for i := start; i < len(candidates); i++ {

		if candidates[i] == target {
			tmp := []int{}
			for j := 0; j < len(*cur); j++ {
				tmp = append(tmp, (*cur)[j])
			}
			tmp = append(tmp, candidates[i])
			*re = append(*re, tmp)
			if len(*cur) > 0 {
				*cur = (*cur)[0 : len(*cur)-1]
			}
			return
		} else if candidates[i] < target {
			*cur = append(*cur, candidates[i])
			combinationSumCore(candidates, i, target-candidates[i], cur, re)
			//*cur = (*cur)[0 : len(*cur)-1]
		} else {
			if len(*cur) > 0 {
				*cur = (*cur)[0 : len(*cur)-1]
			}
			return
		}
	}
	if len(*cur) > 0 {
		*cur = (*cur)[0 : len(*cur)-1]
	}
}*/

func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	backtracking([]int{}, candidates, 0, target, &result)
	return result
}

func backtracking(subset, candidates []int, sum, target int, result *[][]int) {
	if sum == target {
		*result = append(*result, subset)
		return
	} else if sum > target {
		return
	}
	for i := range candidates {
		//创建当前子集合的副本
		//subsetCopy := make([]int, 0)
		var subsetCopy []int
		//	subsetCopy:=[]int{}
		subsetCopy = append(subsetCopy, subset...)
		backtracking(append(subsetCopy, candidates[i]), candidates[i:], sum+candidates[i], target, result)
	}
}

func main() {
	a := []int{2, 3, 5}
	re := combinationSum(a, 8)
	fmt.Println(re)
}

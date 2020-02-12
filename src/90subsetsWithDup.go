package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	if nums == nil {
		return nil
	}
	re := make([][]int, 0)
	sort.Ints(nums)
	helper5(nums, []int{}, &re)
	return re
}

func helper5(nums []int, path []int, re *[][]int) {
	*re = append(*re, append([]int{}, path...))
	for i := 0; i < len(nums); {
		helper5(nums[i+1:], append(path, nums[i]), re)
		i++
		for i < len(nums) && nums[i] == nums[i-1] {
			i++
		}
	}
}

func main() {
	nums := []int{1, 1, 2, 2}
	re := subsetsWithDup(nums)
	fmt.Println(re)
}

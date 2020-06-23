package src

import (
	"fmt"
)

func TwoSum(nums []int, target int) []int {

	for index1, value1 := range nums {
		for index2, value2 := range nums {
			if (index1 != index2) && (value1+value2 == target) {
				return []int{index1, index2}
			}
		}
	}
	return nil

}

func twoSum1(nums []int, target int) []int {
	m := make(map[int]int)
	for i, n := range nums {
		_, prs := m[n] //查看map元素是否存在
		if prs {
			return []int{m[n], i}
		} else {
			m[target-n] = i
		}
	}
	return nil
}
func main() {
	nums := []int{2, 3, 4}
	target := 6
	re := twoSum1(nums, target)
	fmt.Print(re)
}

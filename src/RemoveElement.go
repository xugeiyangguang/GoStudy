package src

import "fmt"

func removeElement(nums []int, val int) int {
	re := 0
	for i := 0; i < len(nums); {
		if val == nums[i] {
			nums = append(nums[:i], nums[i+1:]...)
			continue
		}
		i++
	}
	return len(nums) - re

}
func main() {
	nums := []int{3, 2, 2, 3}
	fmt.Print(removeElement(nums, 2))
}

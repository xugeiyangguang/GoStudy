package src

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)

}
func removeDuplicates(nums []int) int {
	lastIndex := 0
	cur := 1
	for cur < len(nums) {
		if nums[cur] == nums[lastIndex] {
			cur++
		} else {
			nums[lastIndex+1] = nums[cur]
			lastIndex++
			cur++
		}
	}
	return lastIndex + 1
}

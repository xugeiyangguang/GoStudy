package unitTestExcercise

import "math/rand"

type Solution struct {
	arr []int
}

func Constructor2(nums []int) Solution {
	return Solution{arr: append([]int{}, nums...)}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.arr
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	nums := append([]int{}, this.arr...)
	for i := len(nums) - 1; i >= 0; i-- {
		tmp := rand.Intn(i + 1)
		nums[i], nums[tmp] = nums[tmp], nums[i]
	}
	return nums
}

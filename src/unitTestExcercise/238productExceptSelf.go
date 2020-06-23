package unitTestExcercise

func productExceptSelf(nums []int) []int {
	a := make([]int, 0)
	b := make([]int, 0)
	a = append(a, 1)
	b = append(b, 1)
	for i := 1; i < len(nums); i++ {
		a = append(a, a[i-1]*nums[i-1])
		b = append(b, b[i-1]*nums[len(nums)-i])
	}
	c := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		c = append(c, a[i]*b[len(nums)-i-1])
	}
	return c

}

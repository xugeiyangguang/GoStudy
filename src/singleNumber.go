package main

func singleNumber(nums []int) int {
	var re int
	for _, v := range nums {
		re = re ^ v
	}
	return re
}

func main() {
	nums := []int{2, 1, 2}
	println(singleNumber(nums))
}

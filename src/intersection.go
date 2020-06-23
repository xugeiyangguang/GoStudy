package src

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	mapp := make(map[int]int)
	for _, v := range nums1 {
		_, ok := mapp[v]
		if ok {
			mapp[v] = mapp[v] + 1
		} else {
			mapp[v] = 1
		}

	}
	re := []int{}
	for _, v := range nums2 {
		val, ok := mapp[v]
		if ok && val > 0 {
			re = append(re, v)
			mapp[v]--
		}
	}
	return re
}

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	re := intersection(nums1, nums2)
	fmt.Print(re)
}

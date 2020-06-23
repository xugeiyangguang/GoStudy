package unitTestExcercise

import "sort"

func wiggleSort(nums []int) {
	tmp := append([]int{}, nums...)
	sort.Ints(tmp)
	m, n := len(nums)-1, (len(nums)+1)>>1-1
	for i := 0; i < len(nums); i++ {
		if i&1 == 1 {
			nums[i] = tmp[m]
			m--
		} else {
			nums[i] = tmp[n]
			n--
		}
	}

}

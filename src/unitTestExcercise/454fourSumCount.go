package unitTestExcercise

import (
	"fmt"
	"sort"
)

func fourSumCount1(A []int, B []int, C []int, D []int) int {
	sort.Ints(A)
	sort.Ints(B)
	sort.Ints(C)
	sort.Ints(D)
	re := 0
	for i := 0; i < len(A); i++ {
		if B[len(B)-1]+C[len(C)-1]+D[len(D)-1] >= 0-A[i] &&
			B[0]+C[0]+D[0] <= 0-A[i] {
			for j := 0; j < len(B); j++ {
				if C[len(C)-1]+D[len(D)-1] >= 0-(A[i]+B[j]) &&
					C[0]+D[0] <= 0-(A[i]+B[j]) {
					for k := 0; k < len(C); k++ {
						if D[len(D)-1] >= 0-(A[i]+B[j]+C[k]) &&
							D[0] <= 0-(A[i]+B[j]+C[k]) {
							for l := 0; l < len(D); l++ {
								if D[l] == 0-(A[i]+B[j]+C[k]) {
									re++
									fmt.Println(i, j, k, l)
									fmt.Println(A[i], B[j], C[k], D[l])
									fmt.Println("-------------------")
									break
								}
							}
						}
					}
				}
			}
		}
	}
	return re
}

func fourSumCount(A []int, B []int, C []int, D []int) int {
	re := 0
	record := make(map[int]int)
	for _, v1 := range A {
		for _, v2 := range B {
			record[v2+v1]++
		}
	}
	for _, v1 := range C {
		for _, v2 := range D {
			re += record[0-v1-v2]

		}
	}

	return re
}

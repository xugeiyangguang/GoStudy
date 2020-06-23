package unitTestExcercise

func countTriplets(arr []int) int {
	dp := [][]int{}
	record := make(map[int]int, 0)
	for i := 0; i < len(arr); i++ {
		tmp := []int{}
		for j := 0; j < len(arr); j++ {
			tmp = append(tmp, 0)
			if i == j {
				record[arr[i]]++
				tmp[i] = arr[i]
			}
		}
		dp = append(dp, tmp)
	}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			dp[i][j] = dp[i][j-1] ^ arr[j]
			record[dp[i][j]] += 1
		}
	}
	re := 0
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[0]); j++ {
			if record[dp[i][j]] >= 2 {
				for k := i + 1; k < len(dp); k++ {
					for kk := j + 1; kk < len(dp[0]); kk++ {
						if dp[k][kk] == dp[i][j] {
							re++
						}
					}
				}
			}
		}
	}

	//for _, v := range record {
	//	if v != 1 {
	//		re += v * (v - 1) / 2
	//	}
	//}
	return re
}

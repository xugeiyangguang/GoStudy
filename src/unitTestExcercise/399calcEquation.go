package unitTestExcercise

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	start := make(map[string][]int, 0)
	end := make(map[string][]int)
	for i := range equations {
		start[equations[i][0]] = append(start[equations[i][0]], i)
		end[equations[i][1]] = append(end[equations[i][1]], i)
	}
	re := []float64{}
	for i := range queries {

		//not exist
		oneStart := start[queries[i][0]]
		oneEnd := end[queries[i][0]]
		if oneStart == nil && oneEnd == nil {
			re = append(re, -1.0)
			continue
		}
		twoStart := start[queries[i][1]]
		twoEnd := end[queries[i][1]]
		if twoEnd == nil && twoStart == nil {
			re = append(re, -1.0)
			continue
		}
		if queries[i][0] == queries[i][1] {
			re = append(re, 1.0)
			continue
		}
		//本身
		for j := range oneStart {
			for k := range twoEnd {
				if oneStart[j] == twoEnd[k] {
					re = append(re, values[oneStart[j]])
				}
			}
		}
		//倒数
		for j := range oneEnd {
			for k := range twoStart {
				if oneEnd[j] == twoStart[k] {
					re = append(re, 1/values[oneEnd[j]])
				}
			}

		}

		//相乘
		for j := range oneStart {
			for k := range twoEnd {
				if equations[oneStart[j]][1] == equations[twoEnd[k]][0] {
					re = append(re, values[oneStart[j]]*values[twoEnd[k]])
					break
				}
			}
		}

		//两个是分子
		for j := range oneStart {
			for k := range twoStart {
				if equations[oneStart[j]][1] == equations[twoStart[k]][1] {
					re = append(re, values[oneStart[j]]/values[twoStart[k]])
				}
			}
		}

		//两个是分母
		for j := range oneEnd {
			for k := range twoEnd {
				if equations[oneEnd[j]][0] == equations[twoEnd[k]][0] {
					re = append(re, values[twoEnd[k]]/values[oneEnd[j]])
				}
			}
		}

	}
	return re
}

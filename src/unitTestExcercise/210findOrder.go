package unitTestExcercise

func findOrder(numCourses int, prerequisites [][]int) []int {
	re := []int{}
	degree := make(map[int]int)
	lesson := make(map[int][]int)
	for i := 0; i < len(prerequisites); i++ {
		degree[prerequisites[i][0]]++
		lesson[prerequisites[i][1]] = append(lesson[prerequisites[i][1]], prerequisites[i][0])
	}
	for i := 0; i < numCourses; i++ {
		_, ok := degree[i]
		if !ok {
			degree[i] = 0
		}
	}
	for len(degree) != 0 {
		flag := 0
		for k, v := range degree {
			if v == 0 {
				flag = 1
				re = append(re, k)
				tmp := lesson[k]
				for j := range tmp {
					degree[tmp[j]]--
				}
				delete(degree, k)
			}
		}
		if flag == 0 {
			return []int{}
		}
	}

	return re
}

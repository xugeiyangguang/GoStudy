package unitTestExcercise

func canFinish(numCourses int, prerequisites [][]int) bool {
	angleMap := make(map[int]int, numCourses)
	for i := 0; i < numCourses; i++ {
		angleMap[i] = 0
	}
	graphMap := make(map[int][]int, 0)
	for i := 0; i < len(prerequisites); i++ {
		angleMap[prerequisites[i][1]]++
		graphMap[prerequisites[i][0]] = append(graphMap[prerequisites[i][0]], prerequisites[i][1])
	}

	for len(angleMap) != 0 {
		tag := 0
		for k, v := range angleMap {
			if v == 0 {
				tmp := graphMap[k]
				delete(angleMap, k)
				for i := 0; i < len(tmp); i++ {
					angleMap[tmp[i]]--
				}
				tag = 1
			}
		}
		if tag == 0 {
			return false
		}
	}
	return true
}

package main

import "strconv"

/*
1. string 是[]rune类型
2. 注意判断map是否存在某个键的方法
*/
func getHint(secret string, guess string) string {
	myMap := make(map[rune]int, len(secret))
	for _, v := range secret {
		_, ok := myMap[v]
		if ok {
			myMap[v] = myMap[v] + 1
		} else {
			myMap[v] = 1
		}
	}
	bull := 0
	cow := 0
	cowSlice := []byte{}
	for i, v := range guess {
		if guess[i] == secret[i] {
			bull++
			myMap[v]--
		} else {
			if myMap[v] != 0 {
				cowSlice = append(cowSlice, v)
			}
		}
	}
	for _, v := range cowSlice {
		if myMap[v] != 0 {
			myMap[v]--
			cow++
		}
	}
	re := strconv.Itoa(bull) + "A" + strconv.Itoa(cow) + "B"
	return re
}

func main() {
	println(getHint("1123", "0111"))

}

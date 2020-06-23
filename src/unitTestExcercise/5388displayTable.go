package unitTestExcercise

import (
	"sort"
	"strconv"
)

func displayTable(orders [][]string) [][]string {
	food := make(map[string]int)
	record := make(map[string][]string)
	for i := 0; i < len(orders); i++ {
		record[orders[i][1]] = append(record[orders[i][1]], orders[i][2])
		food[orders[i][2]]++
	}
	var key []int
	for k, _ := range record {
		tmp, _ := strconv.Atoi(k)
		key = append(key, tmp)
	}
	sort.Ints(key)
	var foodKey []string
	for k := range food {
		foodKey = append(foodKey, k)
	}
	sort.Strings(foodKey)
	re := [][]string{}
	re = append(re, []string{"Table"})
	for k, _ := range foodKey {
		re[0] = append(re[0], foodKey[k])
	}
	for i := range key {
		t := key[i]
		re = append(re, []string{strconv.Itoa(t)})
		tmp := record[string(strconv.Itoa(t))]
		foodd := make(map[string]int)
		for j := range tmp {
			foodd[tmp[j]]++
		}
		for j := range foodKey {
			re[i+1] = append(re[i+1], strconv.Itoa(foodd[foodKey[j]]))
		}
	}
	return re
}

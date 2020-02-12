package main

import (
	"fmt"
	"sort"
)

type student struct {
	name string
	age  int
}
type stus []student

func (s stus) Len() int {
	return len(s)
}
func (s stus) Less(i, j int) bool {
	return s[i].age < s[j].age
}
func (s stus) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func main() {
	sts := stus{{"xu", 22}, {"yang", 10}, {"haha", 23}}
	fmt.Println("before", sts)
	sort.Sort(sts)
	fmt.Println("after", sts)
	fmt.Println(sort.IsSorted(sts))
	//sort.Sort(sort.Reverse(sts))
	//fmt.Println("reverse:", sts)
	pos := sort.Search(len(sts), func(i int) bool {
		return sts[i].age >= 22
	})
	fmt.Println(pos)
	a := []int{2, 6, 4, 2, 7, 3}
	//	sort.Ints(a)
	sort.Sort(sort.IntSlice(a))
	//	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	fmt.Println(a)
	fmt.Println(sort.SearchInts(a, 4))

}

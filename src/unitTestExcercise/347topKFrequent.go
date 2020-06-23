package unitTestExcercise

import (
	"container/heap"
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	record := make(map[int]int)
	for _, v := range nums {
		record[v]++
	}
	tmp := [][]int{}
	for k, v := range record {
		tmp = append(tmp, []int{k, v})
	}
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i][1]-tmp[j][1] > 0 {
			return true
		}
		return false
	})
	re := []int{}
	for i := 0; i < k; i++ {
		re = append(re, tmp[i][0])
	}
	return re
}

func topKFrequent1(nums []int, k int) []int {
	record := make(map[int]int)
	for _, v := range nums {
		record[v]++
	}
	h := &Nodeheap{}
	topK := min(len(record), k)
	size := 0
	for k, v := range record {
		if size < topK {
			heap.Push(h, &node{k, v})
			size++
		} else {
			if v > (*h)[0].times {
				heap.Pop(h)
				heap.Push(h, &node{k, v})
			}
		}
	}

	re := []int{}
	for i := 0; i < topK; i++ {
		re = append(re, heap.Pop(h).(*node).val)
	}

	return re
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type node struct {
	val   int
	times int
}
type Nodeheap []*node

func (h Nodeheap) Len() int {
	return len(h)
}
func (h Nodeheap) Less(i, j int) bool {
	return h[i].times < h[j].times
}
func (h Nodeheap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *Nodeheap) Push(x interface{}) {
	*h = append(*h, x.(*node))
}
func (h *Nodeheap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

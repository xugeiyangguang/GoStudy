package main

import "fmt"

func leastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}
	var count int
	var maxTask int
	task := map[byte]int{}

	for _, v := range tasks {
		task[v]++
		if task[v] > maxTask {
			count = 1
			maxTask = task[v]
		} else if task[v] == maxTask {
			count++
		}
	}

	idles := maxTask - 1
	idleSize := n - count + 1
	maxTaskSpace := maxTask * count

	return Max(len(tasks), maxTaskSpace+idles*idleSize)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	//	tasks := []byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}
	tasks := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	re := leastInterval(tasks, 2)
	fmt.Println(re)
}

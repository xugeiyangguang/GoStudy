package main

import (
	"Simplemath"
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	a, _ := strconv.Atoi(args[0])
	b, _ := strconv.Atoi(args[1])

	re := Simplemath.Add(a, b)
	fmt.Println("结果是：", re)

}

package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	re := []int{}
	spiralOrderCore(matrix, 0, 0, len(matrix)-1, len(matrix[0])-1, &re)
	return re
}

func spiralOrderCore(matrix [][]int, row int, col int, rows int, cols int, re *[]int) {

	for i := col; i <= cols; i++ {
		*re = append(*re, matrix[row][i]) //在函数里面对切片进行增加，使用的是地址，需要传地址
	}
	for i := row + 1; i <= rows; i++ {
		*re = append(*re, matrix[i][cols])
	}
	for i := cols - 1; i >= col && rows != row; i-- {
		*re = append(*re, matrix[rows][i])
	}
	for i := rows - 1; i > row && col != cols; i-- {
		*re = append(*re, matrix[i][col])
	}
	if rows-row <= 1 || cols-col <= 1 {
		return
	}
	spiralOrderCore(matrix, row+1, col+1, rows-1, cols-1, re)
}

func main() {
	matrix := [][]int{{1, 2, 3}}
	re := spiralOrder(matrix)
	fmt.Println(re)
}

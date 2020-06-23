package unitTestExcercise

func numIslands(grid [][]byte) int {
	if grid == nil {
		return 0
	}
	re := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				re++
				toZero(grid, i, j)
			}
		}
	}
	return re
}

func toZero(grid [][]byte, i, j int) {
	grid[i][j] = '0'
	if i-1 >= 0 && grid[i-1][j] == '1' {
		toZero(grid, i-1, j)
	}
	if i+1 < len(grid) && grid[i+1][j] == '1' {
		toZero(grid, i+1, j)
	}
	if j-1 >= 0 && grid[i][j-1] == '1' {
		toZero(grid, i, j-1)
	}
	if j+1 < len(grid[0]) && grid[i][j+1] == '1' {
		toZero(grid, i, j+1)
	}
}

package main

func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])

	var maxArea int
	var dfs func(row, col int) int
	dfs = func(row, col int) int {
		if row < 0 || row >= m || col < 0 || col >= n || grid[row][col] == 0 {
			return 0
		}
		grid[row][col] = 0
		return 1 + dfs(row-1, col) + dfs(row, col-1) + dfs(row, col+1) + dfs(row+1, col)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				maxArea = max(maxArea, dfs(i, j))
			}
		}
	}
	return maxArea
}

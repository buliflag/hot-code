package main

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	count := 0

	var dfs func(r, c int)
	dfs = func(r, c int) {
		// 判断区间范围和 岛屿述职
		if r < 0 || r >= m || c < 0 || c >= n || grid[r][c] == '0' {
			return
		}
		//让岛屿消失
		grid[r][c] = '0'

		//遍历上下左右
		dfs(r-1, c)
		dfs(r+1, c)
		dfs(r, c-1)
		dfs(r, c+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				count++
				dfs(i, j)
			}
		}
	}
	return count
}

package main

func closedIsland(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])

	var dfs func(row, col int)
	dfs = func(row, col int) {
		if row < 0 || row >= m || col < 0 || col >= n || grid[row][col] == 1 {
			return
		}
		grid[row][col] = 1
		dfs(row-1, col)
		dfs(row+1, col)
		dfs(row, col-1)
		dfs(row, col+1)
	}
	//去除边界线的陆地
	//左右边列
	for i := 0; i < m; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}

	for i := 0; i < n; i++ {
		dfs(0, i)
		dfs(m-1, i)
	}
	count := 0
	//正常走了，边界陆地也被设置为海洋了
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				count++
				dfs(i, j)
			}
		}
	}
	return count

}

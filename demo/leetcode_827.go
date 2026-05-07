package main

func largestIsland(grid [][]int) int {
	n := len(grid)
	//记录不同id的陆地面积
	landIdArea := make(map[int]int)
	isLandId := 2

	var dfs func(r, c, id int) int
	dfs = func(r, c, id int) int {
		//终止条件
		if r < 0 || r >= n || c < 0 || c >= n || grid[r][c] != 1 {
			return 0
		}
		grid[r][c] = id
		return 1 + dfs(r-1, c, id) + dfs(r, c-1, id) + dfs(r, c+1, id) + dfs(r+1, c, id)
	}

	//记录每块陆地面积
	maxArea := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				area := dfs(i, j, isLandId)
				landIdArea[isLandId] = area
				isLandId++
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	//化水为陆后的处理
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				//定义整块大陆连接判断
				connectIsLand := make(map[int]bool)
				dirs := [][]int{
					{i, j + 1},
					{i, j - 1},
					{i - 1, j},
					{i + 1, j},
				}
				curArea := 1
				for _, dir := range dirs {
					r, w := dir[0], dir[1]
					if r >= 0 && r < n && w >= 0 && w < n && grid[r][w] > 1 {
						id := grid[r][w]
						//计算没处理过的新大陆
						if connectIsLand[id] == false {
							curArea += landIdArea[id]
							connectIsLand[id] = true
						}

					}
				}
				if curArea > maxArea {
					maxArea = curArea
				}
			}

		}
	}
	return maxArea

}

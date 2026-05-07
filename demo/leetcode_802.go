package main

func eventualSafeNodes(graph [][]int) []int {
	//初始化，入度表和表关系
	n := len(graph)
	indegress := make([]int, n)

	adjacent := make([][]int, n)

	for i := 0; i < n; i++ {

		for _, node := range graph[i] {
			indegress[i]++
			//
			adjacent[node] = append(adjacent[node], i)
		}
	}

	//queue构建
	queue := []int{}

	for i := 0; i < n; i++ {
		if indegress[i] == 0 {
			queue = append(queue, i)
		}
	}

	//遍历解锁，选出符合条件的
	successNode := []int{}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		for _, next := range adjacent[node] {
			indegress[next]--
			if indegress[next] == 0 {
				queue = append(queue, next)
			}
		}
	}
	for i, d := range indegress {
		if d == 0 {
			successNode = append(successNode, i)
		}
	}
	return successNode

}

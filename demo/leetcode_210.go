package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	//初始化，入度表和 依赖序列
	indegrees := make([]int, numCourses)
	adjacent := make([][]int, numCourses)

	for _, prerequisite := range prerequisites {
		cur, pre := prerequisite[0], prerequisite[1]
		indegrees[cur]++
		adjacent[pre] = append(adjacent[pre], cur)
	}
	//初始队列
	var queue []int
	for i := 0; i < numCourses; i++ {
		if indegrees[i] == 0 {
			queue = append(queue, i)
		}
	}

	//进行解锁，并记录路径
	count := 0
	passRes := []int{}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		count++
		passRes = append(passRes, cur)
		for _, next := range adjacent[cur] {
			//寻找到下一个顺畅的课程
			indegrees[next]--
			if indegrees[next] == 0 {
				queue = append(queue, next)
			}

		}

	}
	if count == numCourses {
		return passRes
	}
	return []int{}

}

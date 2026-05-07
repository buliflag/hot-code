package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	//入度表
	indegrees := make([]int, numCourses)
	adjacent := make([][]int, numCourses)

	//数据准备
	for _, prerequisite := range prerequisites {
		cur, pre := prerequisite[0], prerequisite[1]
		indegrees[cur]++
		adjacent[pre] = append(adjacent[pre], cur)
	}

	//构建可执行队列

	queue := make([]int, numCourses)

	for i := 0; i < numCourses; i++ {
		if indegrees[i] == 0 {
			queue = append(queue, i)
		}
	}
	//上课数量
	count := 0
	for len(queue) > 0 {
		//上课了
		pre := queue[0]
		//出列
		queue = queue[1:]
		count++
		for _, i := range adjacent[pre] {
			//这门课的前置课程完成了
			indegrees[i]--
			if indegrees[i] == 0 {
				//入队
				queue = append(queue, i)
			}
		}

	}
	return count == numCourses
}

func main() {
	numCourses := 2
	prerequisites := make([][]int, 1)
	prerequisites[0] = []int{1, 0}

	print(canFinish(numCourses, prerequisites))

}

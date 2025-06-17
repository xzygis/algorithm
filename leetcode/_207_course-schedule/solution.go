package _207_course_schedule

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make(map[int][]int)  // 存储依赖A的节点有[B,C]
	indegree := make(map[int]int) // 存储每个节点的入度
	for _, pair := range prerequisites {
		graph[pair[1]] = append(graph[pair[1]], pair[0])
		indegree[pair[0]]++
	}

	var q []int
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			q = append(q, i)
		}
	}

	visited := 0
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		visited++
		for _, v := range graph[cur] {
			// 降低依赖该节点的节点的入度
			indegree[v]--
			if indegree[v] == 0 {
				// 入度为0，则表示可以学习了
				q = append(q, v)
			}
		}
	}

	return visited == numCourses
}

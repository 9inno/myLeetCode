package leetcode

func findOrder(numCourses int, prerequisites [][]int) []int {
	inDegree := make([]int, numCourses)
	outEdge := make([][]int, numCourses)
	result := []int{}

	for i:=0; i<len(prerequisites); i++ {
		inDegree[prerequisites[i][0]]++
		outEdge[prerequisites[i][1]] = append(outEdge[prerequisites[i][1]], prerequisites[i][0])
	}

	queue := []int{}

	for i:=0; i<numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	if len(queue) == 0 {
		return result
	}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node)

		for i:=0; i<len(outEdge[node]); i++ {
			out := outEdge[node][i]
			inDegree[out]--
			if inDegree[out] == 0 {
				queue = append(queue, out)
			}
		}
	}

	if len(result) == numCourses {
		return result
	}

	return []int{}
}

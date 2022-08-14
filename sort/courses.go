package sort

// this uses Kahn's algorithm to solve in O(v + e) time/space complexity
// https://www.geeksforgeeks.org/topological-sorting-indegree-based-solution/
func findOrder(numCourses int, prerequisites [][]int) []int {
	adjList := make(map[int][]int)
	indegree := make(map[int]int)

	for _, pair := range prerequisites {
		dest := pair[0]
		src := pair[1]
		adjList[src] = append(adjList[src], dest)

		indegree[dest]++
	}

	zeroIndegreeQueue := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		_, ok := indegree[i]
		if !ok {
			zeroIndegreeQueue = append(zeroIndegreeQueue, i)
		}
	}

	topologicalSortedOrder := make([]int, 0)

	for len(zeroIndegreeQueue) > 0 {
		vertex := zeroIndegreeQueue[0]
		zeroIndegreeQueue = zeroIndegreeQueue[1:]
		topologicalSortedOrder = append(topologicalSortedOrder, vertex)

		if adjacent, ok := adjList[vertex]; ok {
			for _, neighbor := range adjacent {
				indegree[neighbor]--

				if indegree[neighbor] == 0 {
					zeroIndegreeQueue = append(zeroIndegreeQueue, neighbor)
				}
			}
		}
	}

	if len(topologicalSortedOrder) == numCourses {
		return topologicalSortedOrder
	}

	return []int{}
}

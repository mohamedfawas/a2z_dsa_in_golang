package coursescheduleleetcode

func canFinish(numCourses int, prerequisites [][]int) bool {
	// Create an adjacency list and an in-degree array
	adjList := make([][]int, numCourses)
	inDegree := make([]int, numCourses)

	// Build the graph
	for _, prereq := range prerequisites {
		course, pre := prereq[0], prereq[1]
		adjList[pre] = append(adjList[pre], course)
		inDegree[course]++
	}

	// Initialize the queue with courses having in-degree 0
	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// Perform BFS
	visited := 0
	for len(queue) > 0 {
		// Dequeue a course
		current := queue[0]
		queue = queue[1:]
		visited++

		// Decrease the in-degree of its neighbors
		for _, neighbor := range adjList[current] {
			inDegree[neighbor]--
			// If in-degree becomes 0, add to the queue
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// If we visited all courses, return true
	return visited == numCourses
}

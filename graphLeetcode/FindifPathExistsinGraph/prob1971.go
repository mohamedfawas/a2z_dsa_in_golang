package findifpathexistsingraphleetcode

func validPath(n int, edges [][]int, source int, destination int) bool {
	// Create an adjacency list for the graph
	graph := make(map[int][]int)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	// Perform BFS to check if a valid path exists
	visited := make([]bool, n)
	queue := []int{source}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// If we reach the destination, return true
		if current == destination {
			return true
		}

		// Mark the current node as visited
		if visited[current] {
			continue
		}
		visited[current] = true

		// Add all unvisited neighbors to the queue
		for _, neighbor := range graph[current] {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
			}
		}
	}

	// If we exhaust the BFS without reaching the destination, return false
	return false
}

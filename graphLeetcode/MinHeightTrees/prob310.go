package minheighttreesleetcode

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	// Step 1: Build the graph using an adjacency list
	graph := make([][]int, n)
	degree := make([]int, n)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
		degree[a]++
		degree[b]++
	}

	// Step 2: Find all initial leaves
	leaves := []int{}
	for i := 0; i < n; i++ {
		if degree[i] == 1 {
			leaves = append(leaves, i)
		}
	}

	// Step 3: Trim the leaves until 1 or 2 nodes are left
	remainingNodes := n
	for remainingNodes > 2 {
		newLeaves := []int{}
		remainingNodes -= len(leaves)
		for _, leaf := range leaves {
			for _, neighbor := range graph[leaf] {
				degree[neighbor]--
				if degree[neighbor] == 1 {
					newLeaves = append(newLeaves, neighbor)
				}
			}
		}
		leaves = newLeaves
	}

	// Step 4: The remaining nodes are the roots of MHTs
	return leaves
}

package bfsgraphimplementation

import "fmt"

// Graph represents an undirected graph using adjacency list
type Graph struct {
	Vertices int           // Number of vertices in the graph
	AdjList  map[int][]int // Adjacency list to store graph edges
}

// CreateGraph initializes a new graph with given number of vertices
func CreateGrap(vertices int) *Graph {
	return &Graph{
		Vertices: vertices,
		AdjList:  make(map[int][]int),
	}
}

// AddEdge adds an edge between vertex v1 and v2
func (g *Graph) AddEdge(v1, v2 int) {
	// Since this is an undirected graph, we add edges for both directions
	g.AdjList[v1] = append(g.AdjList[v1], v2)
	g.AdjList[v2] = append(g.AdjList[v2], v1)
}

// BFS performs Breadth-First Search starting from the given vertex
func (g *Graph) BFS(startVertex int) {
	// Create a visited map to keep track of visited vertices
	// Using map as a set to store visited vertices
	visited := make(map[int]bool)

	// Create a queue for BFS
	// We'll use a slice as a queue
	queue := []int{}

	// Mark the start vertex as visited and enqueue it
	visited[startVertex] = true
	queue = append(queue, startVertex)

	fmt.Printf("BFS starting from vertex %d :", startVertex)

	// Continue until queue is empty
	for len(queue) > 0 {
		// Dequeue a vertex from queue and print it
		vertex := queue[0] // Get the first element
		queue = queue[1:]  // Remove the first element
		fmt.Printf("%d ", vertex)

		// Get all adjacent vertices of the dequeued vertex
		// For each adjacent vertex that is not visited:
		for _, neighbour := range g.AdjList[vertex] {
			// If this adjacent vertex is not visited
			if !visited[neighbour] {
				// Mark it visited and enqueue it
				visited[neighbour] = true
				queue = append(queue, neighbour)
			}
		}
	}
	fmt.Println()

}

package main

import (
	"fmt"
)

// Graph represents an undirected graph using adjacency list
type Graph struct {
	vertices int           // Number of vertices in the graph
	adjList  map[int][]int // Adjacency list to store graph edges
}

// CreateGraph initializes a new graph with given number of vertices
func CreateGraph(vertices int) *Graph {
	return &Graph{
		vertices: vertices,
		adjList:  make(map[int][]int),
	}
}

// AddEdge adds an edge between vertex v1 and v2
func (g *Graph) AddEdge(v1, v2 int) {
	// Since this is an undirected graph, we add edges for both directions
	g.adjList[v1] = append(g.adjList[v1], v2)
	g.adjList[v2] = append(g.adjList[v2], v1)
}

// DFS performs Depth-First Search starting from the given vertex
func (g *Graph) DFS(startVertex int) {
	// Create a visited map to keep track of visited vertices
	visited := make(map[int]bool)

	// Call the recursive helper function to print DFS traversal
	fmt.Printf("DFS starting from vertex %d: ", startVertex)
	g.dfsRecursive(startVertex, visited)
	fmt.Println()
}

// dfsRecursive is a helper function that implements DFS recursively
func (g *Graph) dfsRecursive(vertex int, visited map[int]bool) {
	// Mark the current vertex as visited and print it
	visited[vertex] = true
	fmt.Printf("%d ", vertex)

	// Recur for all adjacent vertices that have not been visited
	for _, neighbor := range g.adjList[vertex] {
		if !visited[neighbor] {
			g.dfsRecursive(neighbor, visited)
		}
	}
}

func main() {
	// Create a graph with 6 vertices (0 to 5)
	graph := CreateGraph(6)

	// Add edges to create the following graph:
	//    0 --- 1 --- 2
	//    |     |     |
	//    3 --- 4 --- 5
	graph.AddEdge(0, 1)
	graph.AddEdge(1, 2)
	graph.AddEdge(0, 3)
	graph.AddEdge(1, 4)
	graph.AddEdge(2, 5)
	graph.AddEdge(3, 4)
	graph.AddEdge(4, 5)

	// Perform DFS using recursive approach
	graph.DFS(0)

}

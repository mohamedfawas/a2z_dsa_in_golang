package main

import "fmt"

type Graph struct {
	Vertices int
	AdjList  map[int][]int
}

func CreateGraph(vertices int) *Graph {
	return &Graph{

		Vertices: vertices,
		AdjList:  make(map[int][]int),
	}
}

func (g *Graph) AddEdge(v1, v2 int) {
	g.AdjList[v1] = append(g.AdjList[v1], v2)
	g.AdjList[v2] = append(g.AdjList[v2], v1)
}

func (g *Graph) BFS(startVertex int) {
	visited := make(map[int]bool)
	queue := []int{}
	visited[startVertex] = true
	queue = append(queue, startVertex)
	fmt.Printf("BFS starting from vertex %d: ", startVertex)

	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]
		fmt.Printf("%d ", vertex)

		for _, neighbour := range g.AdjList[vertex] {
			if !visited[neighbour] {
				visited[neighbour] = true
				queue = append(queue, neighbour)
			}
		}
	}

	fmt.Println()
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

	// Perform BFS starting from vertex 0
	graph.BFS(0)
}

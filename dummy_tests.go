package main

import "fmt"

type Edge struct {
	To     int
	Weight int
}

type Graph struct {
	AdjList [][]Edge
}

func NewGraph(vertices int) *Graph {
	g := &Graph{make([][]Edge, vertices)}
	for i := range vertices {
		g.AdjList[i] = make([]Edge, 0)
	}

	return g
}

func (g *Graph) AddEdge(from, to, weight int) {
	g.AdjList[from] = append(g.AdjList[from], Edge{To: to, Weight: weight})
}

func (g *Graph) PrintGraph() {
	for i := range g.AdjList {
		for _, edge := range g.AdjList[i] {
			fmt.Printf("edge from : %d connected to : %d with weight : %d ", i, edge.To, edge.Weight)
			fmt.Println()
		}
	}
}

func main() {
	g := NewGraph(4)
	g.AddEdge(0, 1, 5)
	g.AddEdge(0, 2, 3)
	g.AddEdge(0, 3, 1)

	g.PrintGraph()
}

package weightedgraph

import "fmt"

/*
- Start by explaining that you're using an adjacency list representation and why (better for sparse graphs)
- Mention that this implementation is for a directed weighted graph (if you need undirected, you'd add edges in both directions)
- Point out that the time complexity for adding an edge is O(1)
- Mention that you can easily modify this to include more properties in the Edge struct if needed
*/

type Edge struct {
	To     int
	Weight int
}

type Graph struct {
	AdjList [][]Edge
}

// It uses an adjacency list representation, which is memory efficient

func NewGraph(vertices int) *Graph {
	// Make a new graph with space for all our cities
	g := &Graph{
		// Make space for 'vertices' number of cities
		AdjList: make([][]Edge, vertices),
	}
	// For each city, create an empty list of roads
	for i := range g.AdjList {
		g.AdjList[i] = make([]Edge, 0)
	}

	return g
}

// AddEdge is like drawing a new road between two cities
func (g *Graph) AddEdge(from, to, weight int) {
	g.AdjList[from] = append(g.AdjList[from], Edge{To: to, Weight: weight})
}

func (g *Graph) PrintGraph() {
	// Look at each city
	for i := range g.AdjList {
		// Print the city number
		fmt.Printf("city %d is connected to : ", i)
		// Look at all roads leaving from this city
		for _, edge := range g.AdjList[i] {
			fmt.Printf("City %d (distance : %d) ", edge.To, edge.Weight)
		}
		fmt.Println() // Move to next line for the next city
	}
}

func main() {
	// Let's create a small map with 4 cities (numbered 0 to 3)
	graph := NewGraph(4)

	// Now let's add some roads between our cities
	// Remember: AddEdge(from, to, weight)
	graph.AddEdge(0, 1, 10) // Road from city 0 to city 1, distance 10
	graph.AddEdge(0, 2, 5)  // Road from city 0 to city 2, distance 5
	graph.AddEdge(1, 2, 2)  // Road from city 1 to city 2, distance 2
	graph.AddEdge(2, 3, 7)  // Road from city 2 to city 3, distance 7

	// Let's look at our map!
	graph.PrintGraph()

}

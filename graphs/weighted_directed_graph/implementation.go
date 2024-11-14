package weighteddirectedgraph

// Edge represents a connection between two nodes with a weight
type Edge struct {
	From   int
	To     int
	Weight int
}

// Graph represents a weighted directed graph
type Graph struct {
	// Adjacency list to store the edges
	Edges map[int][]Edge
}

// NewGraph creates a new weighted directed graph
func NewGraph() *Graph {
	return &Graph{
		Edges: make(map[int][]Edge),
	}
}

func (g *Graph) AddEdge(from, to, weight int) {
	g.Edges[from] = append(g.Edges[from], Edge{
		From:   from,
		To:     to,
		Weight: weight,
	})
}

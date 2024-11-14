package dijkstrasalgo

// Graph node structure to store vertex and its weight
type Node struct {
	vertex int
	weight int
}

// Priority Queue implementation
type Item struct {
	distance int
	vertex   int
}

type PriorityQueue []Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Item))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// // Function to find shortest path from source to all vertices
// func ShortestPath(adj [][]Node, V int, src int) []int {
// 	// Create distance array and initialize all distances as infinity
// 	// except source vertex which is initialized as 0
// 	dist := make([]int, V)
// 	for i := range dist {
// 		dist[i] = int(1e9) // Using 1e9 as infinity
// 	}
// 	dist[src] = 0

// 	// Create a priority queue to store vertices and their distances
// 	// pq will store {distance, vertex}
// 	pq := make(PriorityQueue, 0)

// }

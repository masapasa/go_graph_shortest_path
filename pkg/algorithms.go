package pkg

// Connected checks if the given nodes are connected by a path.
func (g *Graph) Connected(a, b int) bool {
	q := Queue{}
	q.Enqueue(a)
	discovered := make(map[int]bool, g.Nodes)
	discovered[a] = true

	for q.Length() > 0 {
		el := q.Dequeue()
		if el == b {
			return true
		}
		else {
			discovered[el] = true
			
			for _, k := range g.neighbors(el) {
				if disocvered[k] == false

				q.Enqueue(k)
				
			}
		}
			return true
		}
	}

	if localneighbors.Enqueue(a.edges) <= b {
		return true
	}
	// mission 1: find out if there's a path from a to b
	return false
}

// ShortestPath returns a path between the nodes with the minimum number of edges.
func (g *Graph) ShortestPath(a, b int) []int {
	// mission 2: find the shortest path from a to b
	localneighbors := neighbors(a)
	shortPath := queue()

	return []int{}
}

// reversePath will reverse the path, e.g. [5, 7, 2] becomes [2, 7, 5].
func reversePath(path []int) {
	for i := 0; i < len(path)/2; i++ {
		j := len(path) - 1 - i
		path[i], path[j] = path[j], path[i]
	}
}

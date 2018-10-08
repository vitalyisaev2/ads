package graph

import (
	"math"
)

var _ DirectedCyclicGraph = (*defaultDirectedCyclicGraph)(nil)

type defaultDirectedCyclicGraph struct {
	defaultDirectedGraph
}

// DijkstraShortestPathes returns shortest pathes between
// node and all the other nodes in graph
func (g *defaultDirectedCyclicGraph) DijkstraShortestPathes(from Node) (map[Node][]Node, error) {

	// check node existence
	if _, exists := g.nodes[from.ID()]; !exists {
		return nil, errNodeDoesNotExist(from)
	}

	// initialize queue and resulting dict with initial values
	var (
		shortest = make(map[NodeID]EdgeWeight, len(g.nodes))
		queue    = newNodeHeap()
		err      error
	)
	// populate heap-based priority queue
	for _, node := range g.nodes {
		if node.ID() != from.ID() {
			err = queue.insert(node, math.Inf(1))
			shortest[node.ID()] = math.Inf(1)
		} else {
			err = queue.insert(node, 0)
			shortest[node.ID()] = 0
		}
		if err != nil {
			return nil, err
		}
	}

	// initialize mapping for the preceding items, lying on the shortest path
	pred := make(map[NodeID]NodeID)

	// on every iteration obtain item with minimal shortest[nodeID] value,
	// than perform relaxation procedure
	for queue.size() != 0 {
		//fmt.Println("QUEUE", queue.String())
		curr, weight := queue.min()
		shortest[curr.ID()] = weight
		//fmt.Println("CURR", curr, weight)
		for neighbourID, edges := range g.edges[curr.ID()] {
			if queue.exists(neighbourID) {
				//fmt.Println(curr.ID(), neighbourID)
				for _, edge := range edges {
					// relaxation procedure
					cost := shortest[curr.ID()] + edge.Weight()
					//fmt.Println("COST", cost, shortest[neighbourID])
					if cost < shortest[neighbourID] {
						shortest[neighbourID] = cost
						if err := queue.update(neighbourID, cost); err != nil {
							return nil, err
						}
						pred[neighbourID] = curr.ID()
					}
				}
			}
		}
	}

	// building results
	//fmt.Println(shortest)
	//fmt.Println(pred)
	results := make(map[Node][]Node)
	for _, node := range g.nodes {
		if node.ID() != from.ID() && !emptyNodeID(pred[node.ID()]) {
			var path []Node
			results[node] = path

			currID := node.ID()
			for currID != from.ID() {
				path = append(path, g.nodes[currID])
				currID = pred[currID]
			}
		}
	}
	return results, nil
}

// NewDirectedCyclicGraph returns DirectedCyclicGraph
func NewDirectedCyclicGraph() DirectedCyclicGraph {
	return &defaultDirectedCyclicGraph{
		defaultDirectedGraph: newDirectedGraph(),
	}
}

package graph

import "math"

var _ DirectedCyclicGraph = (*defaultDirectedGraph)(nil)

type defaultDirectedCyclicGraph struct {
	defaultDirectedGraph
}

// DijkstraShortestPathes returns shortest pathes between
func (g *defaultDirectedCyclicGraph) DijkstraShortestPathes(from Node) (map[Node][]Node, error) {

	// check node existence
	if _, exists := g.nodes[from.ID()]; !exists {
		return nil, errNodeDoesNotExist(from)
	}

	// initialize shortest path with initial values
	shortest := make(map[NodeID]EdgeWeight, len(g.nodes))
	for _, node := range g.nodes {
		if node.ID() != from.ID() {
			shortest[node.ID()] = math.Inf(1)
		} else {
			shortest[node.ID()] = 0
		}
	}

	// initialize mapping for the preceding nodes, lying on the shortest path
	pred := make(map[NodeID]NodeID)

	// for node in Q:
}

// NewDirectedCyclicGraph returns DirectedCyclicGraph
func NewDirectedCyclicGraph() DirectedCyclicGraph {
	return &defaultDirectedCyclicGraph{
		defaultDirectedGraph: newDirectedGraph(),
	}
}

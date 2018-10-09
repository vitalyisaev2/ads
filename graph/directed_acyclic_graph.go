package graph

import (
	"fmt"
	"math"
)

var _ DirectedAcyclicGraph = (*defaultDirectedAcyclicGraph)(nil)

type defaultDirectedAcyclicGraph struct {
	defaultDirectedGraph
}

func (g *defaultDirectedAcyclicGraph) TopologicalSort() ([]Node, error) {

	// estimate amount of incoming degrees for every node
	incomingDegrees := make(map[NodeID]int, g.TotalNodes())
	for nodeID := range g.nodes {
		incomingDegrees[nodeID] = 0
	}
	for _, children := range g.edges {
		for childID := range children {
			incomingDegrees[childID]++
		}
	}

	// look up for the items with zero incoming degree: they are "roots" of the graph
	var next []NodeID
	for nodeID, count := range incomingDegrees {
		if count == 0 {
			next = append(next, nodeID)
		}
	}
	if len(next) == 0 {
		return nil, fmt.Errorf("cyclic graph")
	}

	// append to result every node that has no incoming edges
	var results []Node
	for len(next) != 0 {
		parentID := next[0]
		results = append(results, g.nodes[parentID])
		for childID := range g.edges[parentID] {
			incomingDegrees[childID]--
			if incomingDegrees[childID] == 0 {
				next = append(next, childID)
			}
		}
		next = next[1:]
	}

	return results, nil
}

func (g *defaultDirectedAcyclicGraph) ShortestPath(from, to Node) ([]Node, error) {
	// start/end items must exist in graph
	if _, exists := g.nodes[from.ID()]; !exists {
		return nil, errNodeDoesNotExist(from)
	}
	if _, exists := g.nodes[to.ID()]; !exists {
		return nil, errNodeDoesNotExist(to)
	}

	// obtain ordered list of graph items
	sorted, err := g.TopologicalSort()
	if err != nil {
		return nil, err
	}

	// for every node != from, provide the sum of weights on the shortest path on (from; node)
	shortest := make(map[NodeID]EdgeWeight)
	for nodeID := range g.nodes {
		if nodeID == from.ID() {
			shortest[nodeID] = 0
		} else {
			shortest[nodeID] = math.Inf(1)
		}
	}

	// contains ID of the node preceding the current node on some of the shortest paths
	pred := make(map[NodeID]NodeID)

	// for every node from topological order, take neighbour node such that edge (node, neighbour) exists
	for _, node := range sorted {
		for neighbourID, edges := range g.edges[node.ID()] {

			// determine the (node, neighbour) edge with minimal weight
			var (
				minEdgeWeight = EdgeWeight(math.MaxFloat64)
				minEdgeID     EdgeID
			)
			for _, edge := range edges {
				if minEdgeWeight > edge.Weight() {
					minEdgeWeight = edge.Weight()
					minEdgeID = edge.ID()
				}
			}

			// relaxing procedure
			alternative := shortest[node.ID()] + g.edges[node.ID()][neighbourID][minEdgeID].Weight()
			if alternative < shortest[neighbourID] {
				shortest[neighbourID] = alternative
				pred[neighbourID] = node.ID()
			}
		}
	}

	// build reversed shortest path (to, from)
	var (
		path    = []Node{to}
		childID = to.ID()
	)
	for {
		parentID, exists := pred[childID]
		if !exists {
			return nil, fmt.Errorf("no path exists between (%v, %v)", from.ID(), to.ID())
		}
		path = append(path, g.nodes[parentID])
		if parentID == from.ID() {
			break
		}
		childID = parentID
	}

	// reverse result
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path, nil
}

// NewDirectedAcyclicGraph initializes new DirectedAcyclicGraph
func NewDirectedAcyclicGraph() DirectedAcyclicGraph {
	return &defaultDirectedAcyclicGraph{
		defaultDirectedGraph: newDirectedGraph(),
	}
}

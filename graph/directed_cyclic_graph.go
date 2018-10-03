package graph

import (
	"fmt"
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

	// initialize shortest path with initial values
	shortest := make(map[NodeID]EdgeWeight, len(g.nodes))
	for _, node := range g.nodes {
		if node.ID() != from.ID() {
			shortest[node.ID()] = math.Inf(1)
		} else {
			shortest[node.ID()] = 0
		}
	}

	// initialize mapping for the preceding items, lying on the shortest path
	pred := make(map[NodeID]NodeID)

	// populate heap-based priority queue
	queue := newNodeHeap()
	for _, node := range g.nodes {
		if err := queue.insert(node, shortest[node.ID()]); err != nil {
			return nil, err
		}
	}

	// on every iteration obtain item with minimal shortest[nodeID] value,
	// than perform relaxation procedure
	for queue.size() != 0 {
		curr, _ := queue.min()
		for neighbourID, edges := range g.edges[curr.ID()] {
			if queue.exists(neighbourID) {
				fmt.Println(curr.ID(), neighbourID)
				for _, edge := range edges {
					// relaxation procedure
					v := shortest[curr.ID()] + edge.Weight()
					if v < shortest[neighbourID] {
						shortest[neighbourID] = v
						if err := queue.update(neighbourID, v); err != nil {
							return nil, err
						}
						pred[neighbourID] = curr.ID()
					}
				}
			}
		}
	}

	// building results
	fmt.Println(shortest)
	fmt.Println(pred)
	results := make(map[Node][]Node)
	//fmt.Println(pred)
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

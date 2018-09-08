package graph

import "fmt"

var _ DirectedGraph = (*defaultDirectedGraph)(nil)

type defaultDirectedGraph struct {
	nodes map[NodeID]Node                       // node ID <-> node
	edges map[NodeID]map[NodeID]map[EdgeID]Edge // from node ID <-> to node ID <-> edges (like adjacency list)
}

func (g *defaultDirectedGraph) AddNode(node Node) error {
	if _, exists := g.nodes[node.ID()]; exists {
		return fmt.Errorf("node '%s' already exists", node.ID())
	}
	g.nodes[node.ID()] = node
	return nil
}

func (g *defaultDirectedGraph) RemoveNode(node Node) error {
	if _, exists := g.nodes[node.ID()]; !exists {
		return errNodeDoesNotExist(node)
	}
	delete(g.nodes, node.ID())
	return nil
}

func (g *defaultDirectedGraph) AddEdge(edge Edge, from, to Node) error {
	if _, exists := g.nodes[from.ID()]; !exists {
		return errNodeDoesNotExist(from)
	}
	if _, exists := g.nodes[to.ID()]; !exists {
		return errNodeDoesNotExist(to)
	}

	neighbours, exists := g.edges[from.ID()]
	if !exists {
		neighbours = map[NodeID]map[EdgeID]Edge{}
		g.edges[from.ID()] = neighbours
	}

	edges, exists := neighbours[to.ID()]
	if !exists {
		edges = map[EdgeID]Edge{}
		neighbours[to.ID()] = edges
	}

	if _, exists := edges[edge.ID()]; exists {
		return fmt.Errorf("edge '%s' ('%s', '%s') already exists", edge.ID(), from.ID(), to.ID())
	}

	edges[edge.ID()] = edge
	return nil
}

func (g *defaultDirectedGraph) RemoveEdge(edge Edge, from, to Node) error {
	neighbours, exists := g.edges[from.ID()]
	if !exists {
		return errEdgeDoesNotExist(edge, from, to)
	}
	edges, exists := neighbours[to.ID()]
	if !exists {
		return errEdgeDoesNotExist(edge, from, to)
	}
	_, exists = edges[edge.ID()]
	if !exists {
		return errEdgeDoesNotExist(edge, from, to)
	}

	delete(edges, edge.ID())
	return nil
}

func (g *defaultDirectedGraph) TotalNodes() int {
	return len(g.nodes)
}

func (g *defaultDirectedGraph) TotalEdges() int {
	count := 0
	for _, neighbours := range g.edges {
		for _, edges := range neighbours {
			count += len(edges)
		}
	}
	return count
}

func errNodeDoesNotExist(node Node) error {
	return fmt.Errorf("node '%s' doesn't exist", node.ID())
}

func errEdgeDoesNotExist(edge Edge, from, to Node) error {
	return fmt.Errorf("edge '%s' ('%s', '%s') doesn't exist", edge.ID(), from.ID(), to.ID())
}

func (g *defaultDirectedGraph) TopologicalSort() ([]Node, error) {

	// estimate amount of incoming degrees for every node
	incomingDegrees := make(map[NodeID]int, g.TotalNodes())
	for nodeID := range g.nodes {
		incomingDegrees[nodeID] = 0
	}
	for _, toNodeSet := range g.edges {
		for toNode := range toNodeSet {
			incomingDegrees[toNode]++
		}
	}

	// look up for the nodes with zero incoming degree: they are "roots" of the graph
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

// NewDirectedGraph initializes a directed graph
func NewDirectedGraph() DirectedGraph {
	return &defaultDirectedGraph{
		nodes: make(map[NodeID]Node),
		edges: make(map[NodeID]map[NodeID]map[EdgeID]Edge),
	}
}

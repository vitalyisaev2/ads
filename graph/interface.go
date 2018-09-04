package graph

// NodeID represents node identifier
type NodeID = string

// Node is an arbitrary data type; every node should have unique identifier
// to make different nodes distinguishable within a graph
type Node interface {
	// ID returns node unique ID (in the scope of the graph)
	ID() NodeID
}

// EdgeID represents edge identifier
type EdgeID = string

// Edge is an arbitrary data type; there can be multiple edges between two nodes,
// so every edge should provide unique identifier
type Edge interface {
	// ID returns edge unique ID (in the scope of node pair)
	ID() EdgeID
}

// DirectedGraph describes generic directed graph interface
type DirectedGraph interface {
	// AddNode adds new node to a graph
	AddNode(node Node) error
	// RemoveNode removes node from a graph
	RemoveNode(node Node) error
	// AddEdge adds new edge between two nodes
	AddEdge(edge Edge, from, to Node) error
	// RemoveEdge removes edge between two nodes
	RemoveEdge(edge Edge, from, to Node) error
	// TotalNodes returns the amount of nodes in the graph
	TotalNodes() int
	// TotalEdges returns the amount of edges in the graph
	TotalEdges() int
}

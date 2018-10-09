package graph

import "github.com/stretchr/testify/mock"

// DirectedGraphMock mocks Graph interface
type DirectedGraphMock struct {
	mock.Mock
}

var _ DirectedGraph = (*DirectedGraphMock)(nil)

// AddNode adds new node to a graph
func (m *DirectedGraphMock) AddNode(node Node) error {
	args := m.Called(node)
	return args.Error(0)
}

// RemoveNode removes node from a graph
func (m *DirectedGraphMock) RemoveNode(node Node) error {
	args := m.Called(node)
	return args.Error(0)
}

// AddEdge adds new edge between two items
func (m *DirectedGraphMock) AddEdge(edge Edge, from, to Node) error {
	args := m.Called(edge, from, to)
	return args.Error(0)
}

// RemoveEdge removes edge between two items
func (m *DirectedGraphMock) RemoveEdge(edge Edge, from, to Node) error {
	args := m.Called(edge, from, to)
	return args.Error(0)
}

// TopologicalSort returns ordered list of items such that
// every U stands before V if there is (U, V) edge
func (m *DirectedGraphMock) TopologicalSort() ([]Node, error) {
	args := m.Called()
	return args.Get(0).([]Node), args.Error(1)
}

// ShortestPath returns the shortest path between two items, if there are any
func (m *DirectedGraphMock) ShortestPath(from, to Node) ([]Node, error) {
	args := m.Called()
	return args.Get(0).([]Node), args.Error(1)
}

// TotalNodes returns the amount of items in the graph
func (m *DirectedGraphMock) TotalNodes() int { return m.Called().Int(0) }

// TotalEdges returns the amount of edges in the graph
func (m *DirectedGraphMock) TotalEdges() int { return m.Called().Int(0) }

// NodeMock mocks Node interface
type NodeMock struct {
	mock.Mock
}

// ID returns node unique ID (in the scope of the graph)
func (m *NodeMock) ID() NodeID { return m.Called().Get(0).(NodeID) }

// EdgeMock mocks Edge interface
type EdgeMock struct {
	mock.Mock
}

// ID returns edge unique ID (in the scope of node pair)
func (m *EdgeMock) ID() EdgeID { return m.Called().Get(0).(EdgeID) }

// Weight returns edge's weight
func (m *EdgeMock) Weight() float64 { return m.Called().Get(0).(float64) }

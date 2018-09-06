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

// AddEdge adds new edge between two nodes
func (m *DirectedGraphMock) AddEdge(edge Edge, from, to Node) error {
	args := m.Called(edge, from, to)
	return args.Error(0)
}

// RemoveEdge removes edge between two nodes
func (m *DirectedGraphMock) RemoveEdge(edge Edge, from, to Node) error {
	args := m.Called(edge, from, to)
	return args.Error(0)
}

// TotalNodes returns the amount of nodes in the graph
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

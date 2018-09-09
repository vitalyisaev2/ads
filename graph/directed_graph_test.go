package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type stubNode string

func (node stubNode) ID() NodeID { return NodeID(node) }

type stubEdge struct {
	id     EdgeID
	weight Float64
}

func (edge stubEdge) ID() EdgeID { return EdgeID(edge) }

func TestDirectedGraph_BasicOperations(t *testing.T) {
	g := NewDirectedGraph()

	// add node
	node1 := stubNode("node1")
	err := g.AddNode(node1)
	assert.NoError(t, err)

	// can't add same node twice
	err = g.AddNode(node1)
	assert.Error(t, err)

	// add another node
	node2 := stubNode("node2")
	err = g.AddNode(node2)
	assert.NoError(t, err)

	// add edge between node1 and node2
	edge1 := stubEdge("edge1")
	err = g.AddEdge(edge1, node1, node2)
	assert.NoError(t, err)

	// can't add same edge twice
	err = g.AddEdge(edge1, node1, node2)
	assert.Error(t, err)

	// but it's possible to add another edge for the same pair of nodes
	edge2 := stubEdge("edge2")
	err = g.AddEdge(edge2, node1, node2)
	assert.NoError(t, err)

	// check number of edges and nodes
	assert.Equal(t, 2, g.TotalNodes())
	assert.Equal(t, 2, g.TotalEdges())

	// delete nodes
	err = g.RemoveNode(node1)
	assert.NoError(t, err)
	err = g.RemoveNode(node2)
	assert.NoError(t, err)

	// can't delete twice
	err = g.RemoveNode(node2)
	assert.Error(t, err)

	// delete edges
	err = g.RemoveEdge(edge1, node1, node2)
	assert.NoError(t, err)
	err = g.RemoveEdge(edge2, node1, node2)
	assert.NoError(t, err)

	// can't delete twice
	err = g.RemoveEdge(edge2, node1, node2)
	assert.Error(t, err)

	// check number of edges and nodes
	assert.Equal(t, 0, g.TotalNodes())
	assert.Equal(t, 0, g.TotalEdges())
}

var (
	stubNodeA = stubNode("A")
	stubNodeB = stubNode("B")
	stubNodeC = stubNode("C")
	stubNodeD = stubNode("D")
	stubNodeE = stubNode("E")
)

func TestDirectedGraph_TopologicalSort_LinearCase(t *testing.T) {
	// A -> B -> C -> D
	g := NewDirectedGraph()
	g.AddNode(stubNodeA)
	g.AddNode(stubNodeB)
	g.AddNode(stubNodeC)
	g.AddNode(stubNodeD)
	g.AddEdge(stubEdge(""), stubNodeA, stubNodeB)
	g.AddEdge(stubEdge(""), stubNodeB, stubNodeC)
	g.AddEdge(stubEdge(""), stubNodeC, stubNodeD)

	actual, err := g.TopologicalSort()
	assert.NoError(t, err)
	expected := []Node{stubNodeA, stubNodeB, stubNodeC, stubNodeD}
	assert.Equal(t, expected, actual)
}

func TestDirectedGraph_TopologicalSort_Branching(t *testing.T) {

	// A -> B -> C -> E
	//       \-> D -/
	g := NewDirectedGraph()
	g.AddNode(stubNodeA)
	g.AddNode(stubNodeB)
	g.AddNode(stubNodeC)
	g.AddNode(stubNodeD)
	g.AddNode(stubNodeE)
	g.AddEdge(stubEdge(""), stubNodeA, stubNodeB)
	g.AddEdge(stubEdge(""), stubNodeB, stubNodeC)
	g.AddEdge(stubEdge(""), stubNodeB, stubNodeD)
	g.AddEdge(stubEdge(""), stubNodeC, stubNodeE)
	g.AddEdge(stubEdge(""), stubNodeD, stubNodeE)

	actual, err := g.TopologicalSort()
	assert.NoError(t, err)
	expected1 := []Node{stubNodeA, stubNodeB, stubNodeC, stubNodeD, stubNodeE}
	expected2 := []Node{stubNodeA, stubNodeB, stubNodeD, stubNodeC, stubNodeE}
	expected := [][]Node{expected1, expected2}
	assert.Contains(t, expected, actual)
}

func TestDirectedGraph_TopologicalSort_TwoRoots(t *testing.T) {

	// A
	//  \
	//   -> C -> D
	//  /
	// B
	g := NewDirectedGraph()
	g.AddNode(stubNodeA)
	g.AddNode(stubNodeB)
	g.AddNode(stubNodeC)
	g.AddNode(stubNodeD)
	g.AddEdge(stubEdge(""), stubNodeA, stubNodeC)
	g.AddEdge(stubEdge(""), stubNodeB, stubNodeC)
	g.AddEdge(stubEdge(""), stubNodeC, stubNodeD)

	actual, err := g.TopologicalSort()
	assert.NoError(t, err)
	expected1 := []Node{stubNodeA, stubNodeB, stubNodeC, stubNodeD}
	expected2 := []Node{stubNodeB, stubNodeA, stubNodeC, stubNodeD}
	expected := [][]Node{expected1, expected2}
	assert.Contains(t, expected, actual)
}

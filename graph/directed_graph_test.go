package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type stubNode string

func (node stubNode) ID() NodeID { return NodeID(node) }

type stubEdge struct {
	id     EdgeID
	weight EdgeWeight
}

func (edge stubEdge) ID() EdgeID { return EdgeID(edge.id) }

func (edge stubEdge) Weight() EdgeWeight { return edge.weight }

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
	edge1 := stubEdge{id: "edge1"}
	err = g.AddEdge(edge1, node1, node2)
	assert.NoError(t, err)

	// can't add same edge twice
	err = g.AddEdge(edge1, node1, node2)
	assert.Error(t, err)

	// but it's possible to add another edge for the same pair of nodes
	edge2 := stubEdge{id: "edge2"}
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
	stubNodeF = stubNode("F")
)

func TestDirectedGraph_TopologicalSort_LinearCase(t *testing.T) {
	// A -> B -> C -> D
	g := NewDirectedGraph()
	g.AddNode(stubNodeA)
	g.AddNode(stubNodeB)
	g.AddNode(stubNodeC)
	g.AddNode(stubNodeD)
	g.AddEdge(stubEdge{id: "edge0"}, stubNodeA, stubNodeB)
	g.AddEdge(stubEdge{id: "edge1"}, stubNodeB, stubNodeC)
	g.AddEdge(stubEdge{id: "edge2"}, stubNodeC, stubNodeD)

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
	g.AddEdge(stubEdge{id: "edge0"}, stubNodeA, stubNodeB)
	g.AddEdge(stubEdge{id: "edge1"}, stubNodeB, stubNodeC)
	g.AddEdge(stubEdge{id: "edge2"}, stubNodeB, stubNodeD)
	g.AddEdge(stubEdge{id: "edge3"}, stubNodeC, stubNodeE)
	g.AddEdge(stubEdge{id: "edge4"}, stubNodeD, stubNodeE)

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
	g.AddEdge(stubEdge{id: "edge0"}, stubNodeA, stubNodeC)
	g.AddEdge(stubEdge{id: "edge1"}, stubNodeB, stubNodeC)
	g.AddEdge(stubEdge{id: "edge2"}, stubNodeC, stubNodeD)

	actual, err := g.TopologicalSort()
	assert.NoError(t, err)
	expected1 := []Node{stubNodeA, stubNodeB, stubNodeC, stubNodeD}
	expected2 := []Node{stubNodeB, stubNodeA, stubNodeC, stubNodeD}
	expected := [][]Node{expected1, expected2}
	assert.Contains(t, expected, actual)
}

func shortestPathTestGraph() DirectedGraph {
	g := NewDirectedGraph()
	g.AddNode(stubNodeA)
	g.AddNode(stubNodeB)
	g.AddNode(stubNodeC)
	g.AddNode(stubNodeD)
	g.AddNode(stubNodeE)
	g.AddNode(stubNodeF)
	g.AddEdge(stubEdge{weight: 5}, stubNodeA, stubNodeB)
	g.AddEdge(stubEdge{weight: 3}, stubNodeA, stubNodeC)
	g.AddEdge(stubEdge{weight: 2}, stubNodeB, stubNodeC)
	g.AddEdge(stubEdge{weight: 6}, stubNodeB, stubNodeD)
	g.AddEdge(stubEdge{weight: 7}, stubNodeC, stubNodeD)
	g.AddEdge(stubEdge{weight: 4}, stubNodeC, stubNodeE)
	g.AddEdge(stubEdge{weight: 2}, stubNodeC, stubNodeF)
	g.AddEdge(stubEdge{weight: -1}, stubNodeD, stubNodeE)
	g.AddEdge(stubEdge{weight: 1}, stubNodeD, stubNodeF)
	g.AddEdge(stubEdge{weight: -2}, stubNodeE, stubNodeF)
	return g
}

func TestDirectedGraph_ShortestPath_OK(t *testing.T) {
	g := shortestPathTestGraph()
	result, err := g.ShortestPath(stubNodeB, stubNodeF)
	assert.NoError(t, err)
	assert.Equal(t, []Node{stubNodeB, stubNodeD, stubNodeE, stubNodeF}, result)
}

func TestDirectedGraph_ShortestPath_NoPath(t *testing.T) {
	g := shortestPathTestGraph()
	result, err := g.ShortestPath(stubNodeB, stubNodeA)
	assert.Error(t, err)
	assert.Nil(t, result)
}

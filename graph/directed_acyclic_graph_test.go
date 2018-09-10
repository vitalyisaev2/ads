package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirectedGraph_TopologicalSort_LinearCase(t *testing.T) {
	// A -> B -> C -> D
	g := NewDirectedAcyclicGraph()
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
	g := NewDirectedAcyclicGraph()
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
	g := NewDirectedAcyclicGraph()
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

func shortestPathTestDirectedAcyclicGraph() DirectedAcyclicGraph {
	g := NewDirectedAcyclicGraph()
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
	g := shortestPathTestDirectedAcyclicGraph()
	result, err := g.ShortestPath(stubNodeB, stubNodeF)
	assert.NoError(t, err)
	assert.Equal(t, []Node{stubNodeB, stubNodeD, stubNodeE, stubNodeF}, result)
}

func TestDirectedGraph_ShortestPath_NoPath(t *testing.T) {
	g := shortestPathTestDirectedAcyclicGraph()
	result, err := g.ShortestPath(stubNodeB, stubNodeA)
	assert.Error(t, err)
	assert.Nil(t, result)
}

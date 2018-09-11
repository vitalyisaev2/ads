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

var (
	stubNodeA = stubNode("A")
	stubNodeB = stubNode("B")
	stubNodeC = stubNode("C")
	stubNodeD = stubNode("D")
	stubNodeE = stubNode("E")
	stubNodeF = stubNode("F")
	weight1   = EdgeWeight(1)
	weight2   = EdgeWeight(2)
	weight3   = EdgeWeight(3)
	weight4   = EdgeWeight(4)
	weight5   = EdgeWeight(5)
)

func TestDirectedGraph_BasicOperations(t *testing.T) {
	g := newDirectedGraph()

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

	// but it's possible to add another edge for the same pair of items
	edge2 := stubEdge{id: "edge2"}
	err = g.AddEdge(edge2, node1, node2)
	assert.NoError(t, err)

	// check number of edges and items
	assert.Equal(t, 2, g.TotalNodes())
	assert.Equal(t, 2, g.TotalEdges())

	// delete items
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

	// check number of edges and items
	assert.Equal(t, 0, g.TotalNodes())
	assert.Equal(t, 0, g.TotalEdges())
}

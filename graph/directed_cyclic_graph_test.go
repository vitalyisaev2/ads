package graph

import (
	"testing"

	"github.com/vitalyisaev2/testify/assert"
)

var (
	stubNodeS = stubNode("s")
	stubNodeT = stubNode("t")
	stubNodeX = stubNode("x")
	stubNodeY = stubNode("y")
	stubNodeZ = stubNode("z")
)

func TestDirectedCyclicGraph_Dijkstra(t *testing.T) {
	g := NewDirectedCyclicGraph()
	g.AddNode(stubNodeS)
	g.AddNode(stubNodeT)
	g.AddNode(stubNodeX)
	g.AddNode(stubNodeY)
	g.AddNode(stubNodeZ)
	g.AddEdge(stubEdge{weight: 6}, stubNodeS, stubNodeT)
	g.AddEdge(stubEdge{weight: 4}, stubNodeS, stubNodeY)
	g.AddEdge(stubEdge{weight: 3}, stubNodeT, stubNodeX)
	g.AddEdge(stubEdge{weight: 2}, stubNodeT, stubNodeY)
	g.AddEdge(stubEdge{weight: 1}, stubNodeY, stubNodeT)
	g.AddEdge(stubEdge{weight: 9}, stubNodeY, stubNodeX)
	g.AddEdge(stubEdge{weight: 3}, stubNodeY, stubNodeZ)
	g.AddEdge(stubEdge{weight: 4}, stubNodeX, stubNodeZ)
	g.AddEdge(stubEdge{weight: 5}, stubNodeZ, stubNodeX)
	g.AddEdge(stubEdge{weight: 7}, stubNodeZ, stubNodeS)

	result, err := g.DijkstraShortestPathes(stubNodeS)

	assert.Len(t, result, 4)
	assert.Equal(t, result[stubNodeY], []Node{stubNodeY, stubNodeS})
	assert.Equal(t, result[stubNodeT], []Node{stubNodeT, stubNodeY, stubNodeS})
	assert.Equal(t, result[stubNodeX], []Node{stubNodeX, stubNodeT, stubNodeY, stubNodeS})
	assert.Equal(t, result[stubNodeZ], []Node{stubNodeZ, stubNodeY, stubNodeS})

	assert.NoError(t, err)
}

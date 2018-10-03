package graph

import (
	"fmt"
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
	assert.NoError(t, err)
	fmt.Println(result)
}

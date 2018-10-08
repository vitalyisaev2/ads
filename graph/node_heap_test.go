package graph

import (
	"fmt"
	"math"
	"testing"

	"github.com/vitalyisaev2/testify/assert"
)

func TestNodeHeap_BasicOperations_V1(t *testing.T) {
	heap := newNodeHeap()

	var err error
	err = heap.insert(stubNodeA, weight1)
	assert.NoError(t, err)
	err = heap.insert(stubNodeB, weight2)
	assert.NoError(t, err)
	err = heap.insert(stubNodeC, weight3)
	assert.NoError(t, err)
	err = heap.insert(stubNodeD, weight4)
	assert.NoError(t, err)
	err = heap.insert(stubNodeE, weight5)
	assert.NoError(t, err)

	var (
		node  Node
		value EdgeWeight
	)
	node, value = heap.min()
	assert.Equal(t, stubNodeA, node)
	assert.Equal(t, weight1, value)
	node, value = heap.min()
	assert.Equal(t, stubNodeB, node)
	assert.Equal(t, weight2, value)
	node, value = heap.min()
	assert.Equal(t, stubNodeC, node)
	assert.Equal(t, weight3, value)
	node, value = heap.min()
	assert.Equal(t, stubNodeD, node)
	assert.Equal(t, weight4, value)
	node, value = heap.min()
	assert.Equal(t, stubNodeE, node)
	assert.Equal(t, weight5, value)
	node, value = heap.min()
	assert.Nil(t, node)
	assert.Zero(t, value)

	assert.Empty(t, heap.items)
	assert.Empty(t, heap.heap)
}

func TestNodeHeap_BasicOperations_V2(t *testing.T) {
	heap := newNodeHeap()

	var err error
	err = heap.insert(stubNodeS, EdgeWeight(0))
	assert.NoError(t, err)
	err = heap.insert(stubNodeZ, EdgeWeight(math.Inf(1)))
	assert.NoError(t, err)
	err = heap.insert(stubNodeT, EdgeWeight(math.Inf(1)))
	assert.NoError(t, err)
	err = heap.insert(stubNodeX, EdgeWeight(math.Inf(1)))
	assert.NoError(t, err)
	err = heap.insert(stubNodeY, EdgeWeight(math.Inf(1)))
	assert.NoError(t, err)

	curr, weight := heap.min()
	assert.Equal(t, stubNodeS, curr)
	assert.Equal(t, float64(0), weight)

	fmt.Println(heap.String())

	err = heap.update(stubNodeT.ID(), EdgeWeight(6))
	assert.NoError(t, err)
	fmt.Println(heap.String())

	err = heap.update(stubNodeY.ID(), EdgeWeight(4))
	assert.NoError(t, err)
	fmt.Println(heap.String())

	curr, weight = heap.min()
	assert.Equal(t, stubNodeY, curr)
	assert.Equal(t, EdgeWeight(4), weight)
	curr, weight = heap.min()
	assert.Equal(t, stubNodeT, curr)
	assert.Equal(t, EdgeWeight(6), weight)
}

package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

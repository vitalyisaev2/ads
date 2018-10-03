package graph

import (
	"fmt"
)

type nodeHeapItem struct {
	node     Node
	weight   EdgeWeight
	position int
}

// nodeHeap is a heap implementation with extra memory costs
// which allows to keep heap items mutable
type nodeHeap struct {
	heap  []*nodeHeapItem
	items map[NodeID]*nodeHeapItem
}

func (h *nodeHeap) insert(node Node, weight EdgeWeight) error {
	if _, exists := h.items[node.ID()]; exists {
		return fmt.Errorf("node already exists in heap")
	}

	item := &nodeHeapItem{
		node:     node,
		weight:   weight,
		position: 0,
	}

	// put new item to the end of heap
	h.items[node.ID()] = item
	h.heap = append(h.heap, item)
	item.position = len(h.heap) - 1

	item.position = h.siftUp(item.position)
	return nil
}

func (h *nodeHeap) exists(nodeID NodeID) bool {
	_, exists := h.items[nodeID]
	return exists
}

func (h *nodeHeap) update(nodeID NodeID, newWeight EdgeWeight) error {
	item, exists := h.items[nodeID]
	if !exists {
		return fmt.Errorf("node '%s' doesn't exist in heap", nodeID)
	}

	// swap weights
	oldWeight := item.weight
	item.weight = newWeight

	// sift heap up / down, depending on new value
	if oldWeight > newWeight {
		item.position = h.siftDown(item.position)
	} else if newWeight > oldWeight {
		item.position = h.siftUp(item.position)
	}

	return nil
}

func (h *nodeHeap) siftDown(i int) int {
	for 2*i+1 < len(h.heap) {
		left := 2 * i
		right := 2*i + 1
		j := left
		if right < len(h.heap) && h.heap[right].weight < h.heap[left].weight {
			j = right
		}
		if h.heap[i].weight <= h.heap[j].weight {
			break
		}
		h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
		i = j
	}
	return i
}

func (h *nodeHeap) siftUp(i int) int {
	for {
		parent := (i - 1) / 2
		if h.heap[i].weight >= h.heap[parent].weight {
			break
		}
		h.heap[i], h.heap[parent] = h.heap[parent], h.heap[i]
		i = parent
	}
	return i
}

func (h *nodeHeap) min() (Node, EdgeWeight) {
	if len(h.heap) == 0 {
		return nil, 0
	}

	// take first
	root := h.heap[0]
	delete(h.items, root.node.ID())

	// swap root and last heap element, than sift down
	last := h.heap[len(h.heap)-1]
	h.heap[0] = h.heap[len(h.heap)-1]
	h.heap = h.heap[:len(h.heap)-1]

	// sift down heap and store new position of the element that was last
	last.position = h.siftDown(0)

	return root.node, root.weight
}

func (h *nodeHeap) size() int { return len(h.heap) }

// TODO: may be add constructor that builds heap for O(N): now it's O(N*logN)

func newNodeHeap() *nodeHeap {
	return &nodeHeap{
		items: make(map[NodeID]*nodeHeapItem),
	}
}

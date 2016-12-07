package structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// binaryTreeInt is a tree payload
type binaryTreeInt struct {
	value int
}

func (thisItem *binaryTreeInt) Less(thatItem interface{}) bool {
	return thisItem.value < thatItem.(*binaryTreeInt).value
}

func (thisItem *binaryTreeInt) Equal(thatItem interface{}) bool {
	return thisItem.value == thatItem.(*binaryTreeInt).value
}

// Helper function to compare two trees
func treesEqual(tree1 BinarySearchTree, tree2 BinarySearchTree) bool {
	if tree1.Len() != tree2.Len() {
		return false
	}

	tree1Items := tree1.Items()
	tree2Items := tree2.Items()

	for i := 0; i < len(tree1Items); i++ {
		if !tree1Items[i].Item().Equal(tree2Items[i].Item()) {
			return false
		}
	}
	return true
}

func newBinarySearchTreeFromSlice(items []*binaryTreeInt) (BinarySearchTree, error) {
	tree := NewBinarySearchTree()
	for i := range items {
		err := tree.Insert(items[i])
		if err != nil {
			return nil, err
		}
	}
	return tree, nil
}

func TestBinaryTreeBasic(t *testing.T) {
	var (
		err     error
		tree    BinarySearchTree
		node    BinarySearchTreeNode
		payload []binaryTreeInt
	)

	// Create and populate tree
	tree = NewBinarySearchTree()
	payload = []binaryTreeInt{
		binaryTreeInt{2},
		binaryTreeInt{1},
		binaryTreeInt{7},
		binaryTreeInt{4},
		binaryTreeInt{8},
		binaryTreeInt{3},
		binaryTreeInt{6},
		binaryTreeInt{5},
	}
	for i := range payload {
		err = tree.Insert(&payload[i])
		assert.Nil(t, err)
	}

	// Basic checks
	assert.Equal(t, tree.Len(), 8)
	assert.Equal(t, 2, tree.Root().Item().(*binaryTreeInt).value)
	assert.Equal(t, 1, tree.Min().Item().(*binaryTreeInt).value)
	assert.Equal(t, 8, tree.Max().Item().(*binaryTreeInt).value)

	// Insert again --> error
	err = tree.Insert(&binaryTreeInt{1})
	assert.Error(t, err)

	// Search nodes
	node = tree.Search(&binaryTreeInt{1})
	assert.NotNil(t, node)
	node = tree.Search(&binaryTreeInt{10})
	assert.Nil(t, node)
}

func TestBinaryTreeRemoveLeafWithNoChildren(t *testing.T) {
	var (
		treeInitial     BinarySearchTree
		treeLeafRemoved BinarySearchTree
		err             error
	)

	payloadInitial := []*binaryTreeInt{
		&binaryTreeInt{2},
		&binaryTreeInt{1},
		&binaryTreeInt{7},
		&binaryTreeInt{4},
		&binaryTreeInt{8},
		&binaryTreeInt{3},
		&binaryTreeInt{6},
		&binaryTreeInt{5},
	}

	payloadLeafRemoved := []*binaryTreeInt{
		&binaryTreeInt{2},
		&binaryTreeInt{1},
		&binaryTreeInt{7},
		&binaryTreeInt{4},
		&binaryTreeInt{8},
		&binaryTreeInt{6},
		&binaryTreeInt{5},
	}

	treeInitial, _ = newBinarySearchTreeFromSlice(payloadInitial)
	assert.NotNil(t, treeInitial)
	treeLeafRemoved, _ = newBinarySearchTreeFromSlice(payloadLeafRemoved)
	assert.NotNil(t, treeLeafRemoved)

	err = treeInitial.Remove(&binaryTreeInt{3})
	assert.Nil(t, err)
	assert.True(t, treesEqual(treeInitial, treeLeafRemoved))
}

func TestBinaryTreeRemoveLeafWithOneChild(t *testing.T) {
	var (
		treeInitial     BinarySearchTree
		treeLeafRemoved BinarySearchTree
		err             error
	)

	payloadInitial := []*binaryTreeInt{
		&binaryTreeInt{2},
		&binaryTreeInt{1},
		&binaryTreeInt{7},
		&binaryTreeInt{4},
		&binaryTreeInt{8},
		&binaryTreeInt{3},
		&binaryTreeInt{6},
		&binaryTreeInt{5},
	}

	payloadLeafRemoved := []*binaryTreeInt{
		&binaryTreeInt{2},
		&binaryTreeInt{1},
		&binaryTreeInt{7},
		&binaryTreeInt{4},
		&binaryTreeInt{8},
		&binaryTreeInt{3},
		&binaryTreeInt{5},
	}

	treeInitial, _ = newBinarySearchTreeFromSlice(payloadInitial)
	assert.NotNil(t, treeInitial)
	treeLeafRemoved, _ = newBinarySearchTreeFromSlice(payloadLeafRemoved)
	assert.NotNil(t, treeLeafRemoved)

	err = treeInitial.Remove(&binaryTreeInt{6})
	assert.Nil(t, err)

	assert.True(t, treesEqual(treeInitial, treeLeafRemoved))
}

func TestBinaryTreeRemoveLeafWithTwoChildren(t *testing.T) {
	var (
		treeInitial     BinarySearchTree
		treeLeafRemoved BinarySearchTree
		err             error
	)

	payloadInitial := []*binaryTreeInt{
		&binaryTreeInt{2},
		&binaryTreeInt{1},
		&binaryTreeInt{7},
		&binaryTreeInt{4},
		&binaryTreeInt{8},
		&binaryTreeInt{3},
		&binaryTreeInt{6},
		&binaryTreeInt{5},
	}

	payloadLeafRemoved := []*binaryTreeInt{
		&binaryTreeInt{2},
		&binaryTreeInt{1},
		&binaryTreeInt{7},
		&binaryTreeInt{5},
		&binaryTreeInt{8},
		&binaryTreeInt{3},
		&binaryTreeInt{6},
	}

	treeInitial, _ = newBinarySearchTreeFromSlice(payloadInitial)
	assert.NotNil(t, treeInitial)
	treeLeafRemoved, _ = newBinarySearchTreeFromSlice(payloadLeafRemoved)
	assert.NotNil(t, treeLeafRemoved)

	err = treeInitial.Remove(&binaryTreeInt{4})
	assert.Nil(t, err)

	assert.True(t, treesEqual(treeInitial, treeLeafRemoved))
}

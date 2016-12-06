package structures

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type binaryTreeInt struct {
	value int
}

func (thisItem *binaryTreeInt) Less(thatItem interface{}) bool {
	return thisItem.value < thatItem.(*binaryTreeInt).value
}

func (thisItem *binaryTreeInt) Equal(thatItem interface{}) bool {
	return thisItem.value == thatItem.(*binaryTreeInt).value
}

func newBinaryTreeIntSlice() []binaryTreeInt {
	// Create data and shuffle it
	payload := []binaryTreeInt{
		binaryTreeInt{2},
		binaryTreeInt{1},
		binaryTreeInt{7},
		binaryTreeInt{4},
		binaryTreeInt{8},
		binaryTreeInt{3},
		binaryTreeInt{6},
		binaryTreeInt{5},
	}
	return payload
}

func TestBinaryTree(t *testing.T) {
	var (
		err     error
		tree    BinarySearchTree
		node    BinarySearchTreeNode
		payload []binaryTreeInt
	)

	// Create and populate tree
	tree = NewBinarySearchTree()
	payload = newBinaryTreeIntSlice()
	for i := range payload {
		err = tree.Insert(&payload[i])
		assert.Nil(t, err)
	}

	fmt.Println(tree)

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

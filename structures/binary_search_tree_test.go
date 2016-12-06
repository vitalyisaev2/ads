package structures

import (
	//"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type binaryTreeString struct {
	value string
}

func (thisItem *binaryTreeString) Less(thatItem interface{}) bool {
	return thisItem.value < thatItem.(*binaryTreeString).value
}

func (thisItem *binaryTreeString) Equal(thatItem interface{}) bool {
	return thisItem.value == thatItem.(*binaryTreeString).value
}

func TestBinaryTree(t *testing.T) {
	payload := []binaryTreeString{
		binaryTreeString{"a"},
		binaryTreeString{"b"},
		binaryTreeString{"c"},
	}

	// Create and populate tree
	tree := NewBinarySearchTree()
	for i := range payload {
		tree.Insert(&payload[i])
	}
	// Print tree
	//fmt.Println(tree)

	// Basic checks
	assert.Equal(t, tree.Len(), 3)
	assert.Equal(t, "a", tree.Root().(*binaryTreeString).value)
	assert.Equal(t, "a", tree.Min().(*binaryTreeString).value)
	assert.Equal(t, "c", tree.Max().(*binaryTreeString).value)
}

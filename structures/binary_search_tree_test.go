package structures

import (
	"fmt"
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

	tree := NewBinarySearchTree()
	for i := range payload {
		tree.Insert(&payload[i])
	}
	assert.Equal(t, tree.Len(), 3)
	fmt.Println(tree)
	fmt.Println(tree.Items())
	assert.True(t, tree.Min().Equal(&binaryTreeString{"a"}))
	assert.True(t, tree.Root().Equal(&binaryTreeString{"a"}))
}

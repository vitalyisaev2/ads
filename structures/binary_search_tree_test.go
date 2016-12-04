package structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryTree(t *testing.T) {
	var (
		tree BinarySearchTree
	)

	tree = NewBinarySearchTree()
	assert.Nil(t, tree.Min())
}

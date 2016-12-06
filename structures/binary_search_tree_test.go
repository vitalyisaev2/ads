package structures

import (
	"math/rand"
	"testing"
	"time"

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

func newBinaryTreeString() []binaryTreeString {
	// Create data and shuffle it
	payload := []binaryTreeString{
		binaryTreeString{"a"},
		binaryTreeString{"b"},
		binaryTreeString{"c"},
		binaryTreeString{"d"},
		binaryTreeString{"e"},
	}
	rand.Seed(time.Now().UnixNano())
	for i := range payload {
		j := rand.Intn(i + 1)
		payload[i], payload[j] = payload[j], payload[i]
	}
	return payload
}

func TestBinaryTree(t *testing.T) {
	var (
		err     error
		tree    BinarySearchTree
		payload []binaryTreeString
	)

	// Create and populate tree
	tree = NewBinarySearchTree()
	payload = newBinaryTreeString()
	for i := range payload {
		err = tree.Insert(&payload[i])
		assert.Nil(t, err)
	}

	//fmt.Println(tree)

	// Basic checks
	assert.Equal(t, tree.Len(), 5)
	assert.Equal(t, "a", tree.Min().(*binaryTreeString).value)
	assert.Equal(t, "e", tree.Max().(*binaryTreeString).value)

	// True negative
	err = tree.Insert(&binaryTreeString{"a"})
	assert.Error(t, err)
}

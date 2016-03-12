package structures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoublyLinkedList(t *testing.T) {
	var (
		err  error
		item interface{}
	)

	list := NewDoublyLinkedList()

	// Check setters
	list.PushFront(0)
	list.PushBack(2)
	list.Insert(1, 1)

	// Check head/tail accesability
	assert.Equal(t, 0, list.Head())
	assert.Equal(t, 2, list.Tail())

	// Check random access
	item, err = list.Get(1)
	assert.Nil(t, err)
	assert.Equal(t, 1, item.(int))

	//assert.Equal(t, 1, list.Get(1))
	assert.Equal(t, 3, list.Len())
}

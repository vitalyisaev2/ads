package structures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoublyLinkedList(t *testing.T) {
	var (
		err  error
		item interface{}
		list DoublyLinkedList
	)

	list = NewDoublyLinkedList()

	// Check setters
	list.PushFront(0)
	list.PushBack(3)
	list.Insert(1, 1)
	list.Insert(2, 2)

	// Check head/tail accesability
	assert.Equal(t, 0, list.Head())
	assert.Equal(t, 3, list.Tail())

	// Check random access
	item, err = list.Get(1)
	assert.Nil(t, err)
	assert.Equal(t, 1, item.(int))

	// Check length
	assert.Equal(t, 4, list.Len())

	// Check deleters
	item, err = list.Delete(1)
	assert.Nil(t, err)
	assert.Equal(t, 1, item.(int))

	item, err = list.PopFront()
	assert.Nil(t, err)
	assert.Equal(t, 0, item.(int))

	item, err = list.PopBack()
	assert.Nil(t, err)
	assert.Equal(t, 3, item.(int))

	// Check length
	assert.Equal(t, 1, list.Len())

	// Check iterators
	list = NewDoublyLinkedList()
	list.PushBack("a")
	list.PushBack("b")
	list.PushBack("c")

	items := list.Iter()
	assert.Equal(t, "a", (<-items).(string))
	assert.Equal(t, "b", (<-items).(string))
	assert.Equal(t, "c", (<-items).(string))

	ritems := list.RIter()
	assert.Equal(t, "c", (<-ritems).(string))
	assert.Equal(t, "b", (<-ritems).(string))
	assert.Equal(t, "a", (<-ritems).(string))
}

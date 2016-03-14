// Skiena 3.1.2: Linked data structures
// Just a simple thread-safe linked list with minimal amount of goroutines

package structures

import (
	"errors"
	"sync"
)

const (
	back = iota
	front
)

type DoublyLinkedList interface {
	Get(int) (interface{}, error)
	Head() interface{}
	Tail() interface{}
	Len() int

	Insert(interface{}, int) error
	Delete(int) (interface{}, error)
	PushBack(interface{}) error
	PushFront(interface{}) error
	PopBack() (interface{}, error)
	PopFront() (interface{}, error)

	Iter() <-chan interface{}
	RIter() <-chan interface{}
}

type doublyLinkedList struct {
	sync.RWMutex
	head    *doublyLinkedListNode
	tail    *doublyLinkedListNode
	counter int
}

type doublyLinkedListNode struct {
	item interface{}
	prev *doublyLinkedListNode
	next *doublyLinkedListNode
}

type setter func() error
type deleter func() (interface{}, error)

//---------------------- Exported methods -------------------------

func (list *doublyLinkedList) Insert(item interface{}, position int) error {

	// Create new node
	node := &doublyLinkedListNode{item, nil, nil}

	// If list is empty, do the first step
	if list.Len() == 0 {
		insertInitial := func() error {
			list.head = node
			list.tail = node
			list.counter += 1
			return nil
		}
		return list._set(insertInitial)
	}

	if position == 0 {
		insertFront := func() error {
			list.head.prev = node
			node.next = list.head
			list.head = node
			list.counter += 1
			return nil
		}
		return list._set(insertFront)
	}

	if position == list.Len() {
		insertBack := func() error {
			list.tail.next = node
			node.prev = list.tail
			list.tail = node
			list.counter += 1
			return nil
		}
		return list._set(insertBack)
	}

	if 0 < position && position < list.Len() {
		// Get node that lives on the requested position
		existingNode, err := list.getNodeByPosition(position)
		if err != nil {
			return err
		}

		// Emplace new node between two old ones
		insertByPosition := func() error {
			prevNode := existingNode.prev
			existingNode.prev = node
			prevNode.next = node
			node.next = existingNode
			node.prev = prevNode
			list.counter += 1
			return nil
		}
		return list._set(insertByPosition)
	}

	return errors.New("Invalid insertion position: " + string(position))
}

func (list *doublyLinkedList) Delete(position int) (interface{}, error) {

	if list.Len() == 0 {
		return nil, errors.New("Trying to delete node from empty list")
	}

	if position == 0 {
		deleteFront := func() (interface{}, error) {
			item := list.head.item
			newFront := list.head.next
			newFront.prev = nil
			list.counter -= 1
			return item, nil
		}
		return list._delete(deleteFront)
	}

	if position == list.Len() {
		deleteBack := func() (interface{}, error) {
			item := list.tail.item
			newTail := list.tail.prev
			newTail.next = nil
			list.counter -= 1
			return item, nil
		}
		return list._delete(deleteBack)
	}

	if 0 < position && position < list.Len() {
		// Get node that lives on the requested position
		existingNode, err := list.getNodeByPosition(position)
		if err != nil {
			return nil, err
		}

		deleteByPosition := func() (interface{}, error) {
			prevNode := existingNode.prev
			nextNode := existingNode.next
			prevNode.next = nextNode
			nextNode.prev = prevNode
			list.counter -= 1
			return existingNode.item, nil
		}
		return list._delete(deleteByPosition)
	}

	return nil, errors.New("Invalid deletion position: " + string(position))
}

func (list *doublyLinkedList) PushBack(item interface{}) error {
	return list.Insert(item, list.Len())
}

func (list *doublyLinkedList) PushFront(item interface{}) error {
	return list.Insert(item, 0)
}

func (list *doublyLinkedList) PopBack() (interface{}, error) {
	return list.Delete(list.Len())
}

func (list *doublyLinkedList) PopFront() (interface{}, error) {
	return list.Delete(0)
}

func (list *doublyLinkedList) Get(position int) (interface{}, error) {
	node, err := list.getNodeByPosition(position)
	if err != nil {
		return nil, err
	}
	return node.item, nil
}

func (list *doublyLinkedList) Head() interface{} {
	return list.head.item
}

func (list *doublyLinkedList) Tail() interface{} {
	return list.tail.item
}

func (list *doublyLinkedList) Len() int {
	list.RLock()
	defer list.RUnlock()
	return list.counter
}

func (list *doublyLinkedList) Iter() <-chan interface{} {
	items := make(chan interface{})
	nodes := list.iterNodes(back)
	go func() {
		for node := range nodes {
			items <- node.item
		}
		close(items)
	}()
	return items
}

func (list *doublyLinkedList) RIter() <-chan interface{} {
	items := make(chan interface{})
	nodes := list.iterNodes(front)
	go func() {
		for node := range nodes {
			items <- node.item
		}
		close(items)
	}()
	return items
}

//---------------------- Private methods -------------------------

func (list *doublyLinkedList) _set(fn setter) error {
	list.Lock()
	defer list.Unlock()
	return fn()
}

func (list *doublyLinkedList) _delete(fn deleter) (interface{}, error) {
	list.Lock()
	defer list.Unlock()
	return fn()
}

func (list *doublyLinkedList) iterNodes(direction int) <-chan *doublyLinkedListNode {
	ch := make(chan *doublyLinkedListNode, list.Len())

	// Return empty channel if list has no data
	if list.Len() == 0 {
		close(ch)
		return ch
	}

	// Otherwise fill channel with a nodes (taking into account the direction)
	switch direction {

	case back:
		go func() {
			list.RLock()
			defer list.RUnlock()

			current := list.head
			for {
				ch <- current
				if current.next != nil {
					current = current.next
				} else {
					break
				}
			}
			close(ch)
		}()

	case front:
		go func() {

			list.RLock()
			defer list.RUnlock()

			current := list.tail
			for {
				ch <- current
				if current.prev != nil {
					current = current.prev
				} else {
					break
				}
			}
			close(ch)
		}()
	}
	return ch
}

func (list *doublyLinkedList) getNodeByPosition(position int) (*doublyLinkedListNode, error) {
	if position == 0 {
		return list.head, nil
	}

	if position == list.Len() {
		return list.tail, nil
	}

	if position < 0 && (-1)*list.Len() < position {
		return nil, errors.New("Negative indexing is not implemented yet")
	}

	if position >= 0 && position < list.Len() {
		var node *doublyLinkedListNode
		ch := list.iterNodes(back)
		for i := 0; i <= position; i++ {
			node = <-ch
			//log.Println("Iterated: ", node, i)
		}
		return node, nil
	}

	return nil, errors.New("Invalid getting position: " + string(position))
}

// ----------------------- Fabric -------------------------

func NewDoublyLinkedList() DoublyLinkedList {
	return &doublyLinkedList{}
}

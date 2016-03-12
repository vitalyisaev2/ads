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
	//PopBack() interface{}
	//PopFront() interface{}
	Get(int) (interface{}, error)
	Head() interface{}
	Insert(interface{}, int) error
	Len() int
	PushBack(interface{}) error
	PushFront(interface{}) error
	Tail() interface{}
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
		return list._mutate(insertInitial)
	}

	if position == 0 {
		insertFront := func() error {
			list.head.prev = node
			node.next = list.head
			list.head = node
			list.counter += 1
			return nil
		}
		return list._mutate(insertFront)
	}

	if position == list.Len() {
		insertBack := func() error {
			list.tail.next = node
			node.prev = list.tail
			list.tail = node
			list.counter += 1
			return nil
		}
		return list._mutate(insertBack)
	}

	if (-1)*list.Len() < position && position < list.Len() {
		return list.insertIntoPosition(node, position)
	}

	return errors.New("Invalid insertion position: " + string(position))
}

func (list *doublyLinkedList) PushBack(item interface{}) error {
	return list.Insert(item, list.Len())
}

func (list *doublyLinkedList) PushFront(item interface{}) error {
	return list.Insert(item, 0)
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

//---------------------- Private methods -------------------------

type mutator func() error

func (list *doublyLinkedList) _mutate(fn mutator) error {
	list.Lock()
	defer list.Unlock()
	return fn()
}

func (list *doublyLinkedList) insertIntoPosition(newNode *doublyLinkedListNode, position int) error {

	// Get node that lives on the requested position
	existingNode, err := list.getNodeByPosition(position)
	if err != nil {
		return err
	}

	// Emplace new node between two old ones
	emplaceNode := func() error {
		prevNode := existingNode.prev
		existingNode.prev = newNode
		prevNode.next = newNode
		newNode.next = existingNode
		newNode.prev = prevNode
		list.counter += 1
		return nil
	}
	return list._mutate(emplaceNode)

}

func (list *doublyLinkedList) iterNodes(direction int) <-chan *doublyLinkedListNode {
	ch := make(chan *doublyLinkedListNode)

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
		}
		return node, nil
	}

	return nil, errors.New("Invalid getting position: " + string(position))
}

// ----------------------- Fabric -------------------------

func NewDoublyLinkedList() DoublyLinkedList {
	return &doublyLinkedList{}
}

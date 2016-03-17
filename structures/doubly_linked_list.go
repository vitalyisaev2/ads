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

	list.Lock()
	defer list.Unlock()

	// Create new node
	node := &doublyLinkedListNode{item, nil, nil}

	// If list is empty, do the first step
	if list.counter == 0 {
		list.head = node
		list.tail = node
		list.counter += 1
		return nil
	}

	if position == 0 {
		list.head.prev = node
		node.next = list.head
		list.head = node
		list.counter += 1
		return nil
	}

	if position == list.counter || position == -1 {
		list.tail.next = node
		node.prev = list.tail
		list.tail = node
		list.counter += 1
		return nil
	}

	if 0 < position && position < list.counter {
		// Get node that lives on the requested position
		existingNode, err := list.getNodeByPosition(position)
		if err != nil {
			return err
		}

		// Emplace new node between two old ones
		prevNode := existingNode.prev
		existingNode.prev = node
		prevNode.next = node
		node.next = existingNode
		node.prev = prevNode
		list.counter += 1
		return nil
	}

	return errors.New("Invalid insertion position: " + string(position))
}

func (list *doublyLinkedList) Delete(position int) (interface{}, error) {
	list.Lock()
	defer list.Unlock()

	if list.counter == 0 {
		return nil, errors.New("Trying to delete node from empty list")
	}

	if list.counter == 1 {
		item := list.head.item
		list.head = nil
		list.tail = nil
		list.counter -= 1
		return item, nil
	}

	if position == 0 {
		item := list.head.item
		newFront := list.head.next
		newFront.prev = nil
		list.counter -= 1
		return item, nil
	}

	if position == list.counter || position == -1 {
		item := list.tail.item
		newTail := list.tail.prev
		newTail.next = nil
		list.counter -= 1
		return item, nil
	}

	if 0 < position && position < list.counter {
		// Get node that lives on the requested position
		existingNode, err := list.getNodeByPosition(position)
		if err != nil {
			return nil, err
		}

		prevNode := existingNode.prev
		nextNode := existingNode.next
		prevNode.next = nextNode
		nextNode.prev = prevNode
		list.counter -= 1
		return existingNode.item, nil
	}

	return nil, errors.New("Invalid deletion position: " + string(position))
}

func (list *doublyLinkedList) PushBack(item interface{}) error {
	return list.Insert(item, -1)
}

func (list *doublyLinkedList) PushFront(item interface{}) error {
	return list.Insert(item, 0)
}

func (list *doublyLinkedList) PopBack() (interface{}, error) {
	return list.Delete(-1)
}

func (list *doublyLinkedList) PopFront() (interface{}, error) {
	return list.Delete(0)
}

func (list *doublyLinkedList) Get(position int) (interface{}, error) {
	list.RLock()
	defer list.RUnlock()
	node, err := list.getNodeByPosition(position)
	if err != nil {
		return nil, err
	}
	return node.item, nil
}

func (list *doublyLinkedList) Head() interface{} {
	list.RLock()
	defer list.RUnlock()
	return list.head.item
}

func (list *doublyLinkedList) Tail() interface{} {
	list.RLock()
	defer list.RUnlock()
	return list.tail.item
}

func (list *doublyLinkedList) Len() int {
	list.RLock()
	defer list.RUnlock()
	return list.counter
}

func (list *doublyLinkedList) Iter() <-chan interface{} {
	list.RLock()
	defer list.RUnlock()

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
	list.RLock()
	defer list.RUnlock()

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

// The following functions are not protected with locks, because the higher level functions
// should be responsible for that. Otherwise we can run into various race conditions
func (list *doublyLinkedList) iterNodes(direction int) <-chan *doublyLinkedListNode {
	ch := make(chan *doublyLinkedListNode, list.counter)

	// Return empty channel if list has no data
	if list.counter == 0 {
		close(ch)
		return ch
	}

	// Otherwise fill channel with a nodes (taking into account the direction)
	switch direction {

	// Cannot use goroutines for looping over structure due to race conditions
	case back:
		current := list.head
		for {
			ch <- current
			if current.next != nil {
				current = current.next
			} else {
				break
			}
		}

	case front:
		current := list.tail
		for {
			ch <- current
			if current.prev != nil {
				current = current.prev
			} else {
				break
			}
		}
	}

	close(ch)
	return ch
}

func (list *doublyLinkedList) getNodeByPosition(position int) (*doublyLinkedListNode, error) {
	if position == 0 {
		return list.head, nil
	}

	if position == list.counter {
		return list.tail, nil
	}

	if position < 0 && (-1)*list.counter < position {
		return nil, errors.New("Negative indexing is not implemented yet")
	}

	if position >= 0 && position < list.counter {
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

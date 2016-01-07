// Skiena 3.1.2: pointers and linked data structures
package structures

import (
	"errors"
	"fmt"
	"reflect"
)

// Linked List Record can store any data type
type LinkedListElement interface{}

// Linked List Record is a generic element of any Linked List
type LinkedListRecord struct {
	item LinkedListElement
	next *LinkedListRecord
}

// Method providing capability of the recursive item search
func (record *LinkedListRecord) searchRecursive(item LinkedListElement) *LinkedListRecord {

	if record == nil {
		return nil
	}

	if reflect.DeepEqual(record.item, item) {
		return record
	}

	return record.next.searchRecursive(item)
}

// Linked List data structure
type LinkedList struct {
	last   *LinkedListRecord
	length int
}

// Append an item to linked list
func (list *LinkedList) Append(item LinkedListElement) {
	record := &LinkedListRecord{item: item, next: list.last}
	list.last = record
	list.length += 1
}

// Iterate through LinkedList
func (list *LinkedList) Iter() <-chan *LinkedListRecord {
	var ch chan *LinkedListRecord

	// If list contains only nil element, return empty channel
	if list.length == 0 {
		ch = make(chan *LinkedListRecord)
		close(ch)
		return ch
	}

	// Otherwise fill buffered channel with elements
	ch = make(chan *LinkedListRecord, list.length)
	go func() {
		record := list.last
		ch <- record
		for record.next != nil {
			record = record.next
			ch <- record
		}
		close(ch)
	}()

	return ch
}

// Iterative item search
func (list *LinkedList) SearchIterative(item LinkedListElement) *LinkedListRecord {

	if list.last == nil {
		return nil
	}

	for record := range list.Iter() {
		if reflect.DeepEqual(record.next.item, item) {
			return record
		}
	}

	return nil
}

// Recursive item search
func (list *LinkedList) SearchRecursive(item LinkedListElement) *LinkedListRecord {

	if list.last == nil {
		return nil
	}

	return list.last.searchRecursive(item)
}

// Search a predecessor for a record containing specified item
func (list *LinkedList) searchPredecessor(item LinkedListElement) *LinkedListRecord {
	if list.last == nil || list.last.next == nil {
		return nil
	}

	for record := range list.Iter() {
		if reflect.DeepEqual(record.next.item, item) {
			return record
		}
	}

	return nil
}

// Delete an item from linked list
func (list *LinkedList) Delete(item LinkedListElement) error {

	if list.last == nil {
		err := errors.New("Attempting to perform Delete operation on empty list")
		return err
	}

	predecessor := list.searchPredecessor(item)

	// If list contains the only element
	if predecessor == nil && list.last.next == nil {

		// Delete this element if matches
		if reflect.DeepEqual(list.last.item, item) {
			list.last = nil
			list.length -= 1
			if list.length != 0 {
				msg := fmt.Sprintf("LinkedList length is %d after deleting last element", list.length)
				panic(errors.New(msg))
			}

			// Return error otherwise
		} else {
			msg := fmt.Sprintf("Element %v wasn't found", item)
			return errors.New(msg)
		}
	}

	// Otherwise replace predecessor's pointer
	predecessor.next = predecessor.next.next
	list.length -= 1
	return nil
}

// Compare two linked lists
func (initial *LinkedList) EqualTo(compared *LinkedList) bool {

	if initial.length != compared.length {
		return false
	}

	initialValues := initial.Iter()
	comparedValues := compared.Iter()
	for i := 0; i < initial.length; i++ {
		val1 := (<-initialValues).item
		val2 := (<-comparedValues).item
		//fmt.Printf("%v %v\n", val1, val2)
		//fmt.Printf("%v %v\n", reflect.ValueOf(val1).Elem(), reflect.ValueOf(val2).Elem())
		if !reflect.DeepEqual(val1, val2) {
			return false
		}
	}

	return true
}

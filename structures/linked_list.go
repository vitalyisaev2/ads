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

// Extend LinkedList with the elements of another one
// (inspired by Python's list.extend() function)
func (list *LinkedList) Extend(items []LinkedListElement) {
	for _, item := range items {
		list.Append(item)
	}
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
func (list *LinkedList) searchIterative(item LinkedListElement) *LinkedListRecord {

	if list.last == nil {
		return nil
	}

	for record := range list.Iter() {
		if reflect.DeepEqual(record.item, item) {
			return record
		}
	}

	return nil
}

// Recursive item search
func (list *LinkedList) searchRecursive(item LinkedListElement) *LinkedListRecord {

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
		if record.next != nil {
			if reflect.DeepEqual(record.next.item, item) {
				return record
			}
		}
	}

	return nil
}

// Search proxy function
// Best   - searchRecursive vs searchIterative - 13600x faster
// Middle - searchRecursive vs searchIterative - 2.2x faster
// Worst  - searchRecursive vs searchIterative - 1.7x faster
func (list *LinkedList) Search(item LinkedListElement) *LinkedListRecord {
	return list.searchRecursive(item)
}

// Delete an item from linked list
func (list *LinkedList) Delete(item LinkedListElement) error {

	// Error fabric function
	itemNotFound := func() error {
		msg := fmt.Sprintf("Element %#v wasn't found", item)
		return errors.New(msg)
	}

	if list.last == nil {
		err := errors.New("Attempting to perform Delete operation on empty list")
		return err
	}

	predecessor := list.searchPredecessor(item)

	// Predecessor is not found
	if predecessor == nil {

		// Case 1: There are many items in LinkedList, so the item is just not found
		if list.last.next != nil {
			return itemNotFound()
		}
		// Case 2: Probably there is an only item in LinkedList, therefore it hasn't a predecessor
		// Delete this element if matches
		if reflect.DeepEqual(list.last.item, item) {
			list.last = nil
			list.length -= 1
			return nil
		} else {
			// Return itemNotFound otherwise
			return itemNotFound()
		}
	}

	// Predecessor was found: now replace predecessor's pointer
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
		if !reflect.DeepEqual((<-initialValues).item, (<-comparedValues).item) {
			return false
		}
	}

	return true
}

// Print out LinkedList values (only for debug)
func (list *LinkedList) Print() {
	counter := 1
	for record := range list.Iter() {
		fmt.Printf("%d: item: %#v self: %p next: %p\n", counter, record.item, record, record.next)
		counter++
	}
}

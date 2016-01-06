package structures

import (
	//"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedListAppend(t *testing.T) {

	manuallyFilledList := new(LinkedList)
	apiFilledList := new(LinkedList)

	// Fill LinkedList manually
	{
		first_item := 0
		first_record := &LinkedListRecord{
			item: first_item,
			next: nil,
		}
		second_item := 1
		second_record := &LinkedListRecord{
			item: second_item,
			next: first_record,
		}
		manuallyFilledList.last = second_record
	}

	// Fill LinkedList via Append
	{
		items := []int{0, 1}
		for _, item := range items {
			apiFilledList.Append(item)
		}
	}

	// Check the equality of elements
	assert.Equal(t, manuallyFilledList.last.item, apiFilledList.last.item)
	assert.Equal(t, manuallyFilledList.last.next.item, apiFilledList.last.next.item)

	// Check LinkedList.length attribute
	assert.Equal(t, apiFilledList.length, 2)
}

func TestLinkedListIter(t *testing.T) {

	// Fill new LinkedList with 3 string items
	list := new(LinkedList)
	strings := []string{"first", "second", "third"}
	for _, str := range strings {
		list.Append(str)
	}

	// Check that loop will repeat 3 times exactly
	counter := 0
	for _ = range list.Iter() {
		counter++
	}
	assert.Equal(t, counter, 3)

	// Check the order of items (should be reversed)
	var record *LinkedListRecord
	records := list.Iter()
	record = <-records
	assert.Equal(t, record.item, "third")
	record = <-records
	assert.Equal(t, record.item, "second")
	record = <-records
	assert.Equal(t, record.item, "first")
}

func TestLinkedListIterEmpty(t *testing.T) {

	// Create empty LinkedList
	list := new(LinkedList)

	// Check that loop will exit immideately
	counter := 0
	for _ = range list.Iter() {
		counter++
	}
	assert.Equal(t, counter, 0)
}

func TestLinkedListEqualTo(t *testing.T) {

	// Fill up two LinkedLists with the same elements
	initial := new(LinkedList)
	compared := new(LinkedList)
	initialItems := []int{1, 2, 3, 4, 5}
	comparedItems := []int{1, 2, 3, 4, 5}
	for _, item := range initialItems {
		initial.Append(item)
	}
	for _, item := range comparedItems {
		compared.Append(item)
	}

	// Check equality
	var equality bool
	equality = initial.EqualTo(compared)
	assert.True(t, equality)
}

package structures

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Several usefull typedefs and functions
type complexTestStruct struct {
	num int
	str string
}

func pointersComplexStructFabric() []*complexTestStruct {
	slice := []*complexTestStruct{
		&complexTestStruct{1, "a"},
		&complexTestStruct{2, "b"},
		&complexTestStruct{3, "c"},
		&complexTestStruct{4, "d"},
		&complexTestStruct{5, "e"}}
	return slice
}

func valuesComplexStructFabric() []complexTestStruct {
	slice := []complexTestStruct{
		complexTestStruct{1, "a"},
		complexTestStruct{2, "b"},
		complexTestStruct{3, "c"},
		complexTestStruct{4, "d"},
		complexTestStruct{5, "e"}}
	return slice
}

func pointersLinkedListElementFabric() []LinkedListElement {
	slice := []LinkedListElement{
		LinkedListElement(&complexTestStruct{1, "a"}),
		LinkedListElement(&complexTestStruct{2, "b"}),
		LinkedListElement(&complexTestStruct{3, "c"}),
		LinkedListElement(&complexTestStruct{4, "d"}),
		LinkedListElement(&complexTestStruct{5, "e"})}
	return slice
}

func valuesLinkedListElementFabric() []LinkedListElement {
	slice := []LinkedListElement{
		LinkedListElement(complexTestStruct{1, "a"}),
		LinkedListElement(complexTestStruct{2, "b"}),
		LinkedListElement(complexTestStruct{3, "c"}),
		LinkedListElement(complexTestStruct{4, "d"}),
		LinkedListElement(complexTestStruct{5, "e"})}
	return slice
}

// LinkedList.Append
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

	// Fill up two LinkedLists with pointers to structs
	{
		initial := new(LinkedList)
		compared := new(LinkedList)
		initialItems := pointersComplexStructFabric()
		comparedItems := pointersComplexStructFabric()

		for _, item := range initialItems {
			initial.Append(item)
		}
		for _, item := range comparedItems {
			compared.Append(item)
		}

		// Check equality
		equality := initial.EqualTo(compared)
		assert.True(t, equality)
	}

	// Fill up two LinkedLists with structs as is
	{
		initial := new(LinkedList)
		compared := new(LinkedList)
		initialItems := valuesComplexStructFabric()
		comparedItems := valuesComplexStructFabric()

		for _, item := range initialItems {
			initial.Append(item)
		}
		for _, item := range comparedItems {
			compared.Append(item)
		}

		// Check equality
		equality := initial.EqualTo(compared)
		assert.True(t, equality)
	}
}

func TestLinkedListExtend(t *testing.T) {
	// Create LinkedList via .Append()
	appendedList := new(LinkedList)
	appendedItems := pointersLinkedListElementFabric()
	for _, item := range appendedItems {
		appendedList.Append(item)
	}
	assert.Equal(t, 5, appendedList.length)

	// Create LinkedList via .Extend()
	extendedList := new(LinkedList)
	extendedItems := pointersLinkedListElementFabric()
	extendedList.Extend(extendedItems)
	assert.Equal(t, 5, extendedList.length)

	// Compare the results
	equality := appendedList.EqualTo(extendedList)
	assert.True(t, equality)
}

func TestLinkedListSearch(t *testing.T) {

	// Performs tests for different search implementations and use cases
	searchTest := func(searchMethodName string) {

		// Switcher for different search methods
		searchTestExecute := func(list *LinkedList, desiredElement LinkedListElement) *LinkedListRecord {

			var record *LinkedListRecord
			switch searchMethodName {
			case "searchRecursive":
				record = list.searchRecursive(desiredElement)
			case "searchIterative":
				record = list.searchIterative(desiredElement)
			case "Search":
				record = list.Search(desiredElement)
			}
			return record
		}

		// Case 1: search should find existing element
		{
			list := new(LinkedList)
			list.Extend(pointersLinkedListElementFabric())
			record := searchTestExecute(list, LinkedListElement(&complexTestStruct{3, "c"}))
			if assert.NotNil(t, record) {
				assert.Equal(t, record.item, &complexTestStruct{3, "c"})
			}
		}

		// Case 2: search should not find nonexistent element
		{
			list := new(LinkedList)
			list.Extend(pointersLinkedListElementFabric())
			record := searchTestExecute(list, LinkedListElement(&complexTestStruct{6, "f"}))
			assert.Nil(t, record)
		}
	}

	searchTest("searchRecursive")
	searchTest("Search")
}

func TestLinkedListDelete(t *testing.T) {

	// Case 1: delete existing element in "long" LinkedList
	// and check the integrity of LinkedList
	{
		// Construct a LinkedList and delete one element
		testingLinkedList := new(LinkedList)
		testingLinkedList.Extend(pointersLinkedListElementFabric())
		removedElement := LinkedListElement(&complexTestStruct{3, "c"})
		testingLinkedListInitialLength := testingLinkedList.length
		result := testingLinkedList.Delete(removedElement)
		assert.Nil(t, result)
		testingLinkedListResultLength := testingLinkedList.length

		// Construct a LinkedList without a deleted element
		checkingLinkedList := new(LinkedList)
		checkingLinkedListItems := pointersLinkedListElementFabric()
		checkingLinkedListItems = append(checkingLinkedListItems[:2], checkingLinkedListItems[3:]...)
		checkingLinkedList.Extend(checkingLinkedListItems)

		// Check the equality of LinkedLists
		equality := testingLinkedList.EqualTo(checkingLinkedList)
		if !assert.True(t, equality) {
			fmt.Printf("testingLinkedList (length=%d)\n", testingLinkedList.length)
			testingLinkedList.Print()
			fmt.Printf("checkingLinkedList (length=%d)\n", checkingLinkedList.length)
			checkingLinkedList.Print()
		}

		// Check length decreasing
		assert.Equal(t, testingLinkedListInitialLength-testingLinkedListResultLength, 1)

		// Check that predecessor is pointing to the appropriate record now
		predecessorRecord := testingLinkedList.Search(LinkedListElement(&complexTestStruct{4, "d"}))
		followerRecord := testingLinkedList.Search(LinkedListElement(&complexTestStruct{2, "b"}))
		assert.NotNil(t, predecessorRecord, followerRecord)
		if !assert.Equal(t, predecessorRecord.next, followerRecord) {
			fmt.Printf("testingLinkedList (length=%d)\n", testingLinkedList.length)
			testingLinkedList.Print()
		}
	}

	// Case 2: delete existing element in one-element LinkedList
	// and check the integrity of LinkedList
	{
		testingLinkedList := new(LinkedList)
		removedElement := LinkedListElement(&complexTestStruct{6, "f"})
		testingLinkedList.Append(removedElement)

		predecessorRecord := testingLinkedList.searchPredecessor(removedElement)
		assert.Nil(t, predecessorRecord)

		result := testingLinkedList.Delete(removedElement)
		assert.Nil(t, result)

		assert.Equal(t, 0, testingLinkedList.length)
	}

	// Case 3: delete nonexisting element and catch error
	{
		testingLinkedList := new(LinkedList)
		testingLinkedList.Extend(pointersLinkedListElementFabric())
		removedElement := LinkedListElement(&complexTestStruct{6, "f"})

		predecessorRecord := testingLinkedList.searchPredecessor(removedElement)
		assert.Nil(t, predecessorRecord)

		result := testingLinkedList.Delete(removedElement)
		assert.Error(t, result)
	}

	// Case 4: perform delete on empty LinkedList and catch error
	{
		testingLinkedList := new(LinkedList)
		removedElement := LinkedListElement(&complexTestStruct{6, "f"})

		result := testingLinkedList.Delete(removedElement)
		assert.Error(t, result)
	}
}

func TestLinkedListPrint(t *testing.T) {
	// Case 1: print filled LinkedList
	{
		list := new(LinkedList)
		list.Extend(pointersLinkedListElementFabric())
		list.Print()
	}

	// Case 2: print empty LinkedList
	{
		list := new(LinkedList)
		list.Print()
	}
}

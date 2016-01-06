package structures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type NumericalArray []int
type CharacterArray []rune
type StructArray []struct {
	str string
	num int
}

func TestLinkedListAppend(t *testing.T) {

	manuallyFilledList := new(LinkedList)
	apiFilledList := new(LinkedList)

	// Fill list manually
	{
		item := 0
		record := &LinkedListRecord{
			item: item,
			next: nil,
		}
		manuallyFilledList.last = record
	}

	// Fill list with Append
	{
		item := 0
		apiFilledList.Append(item)
	}

	assert.Equal(t, manuallyFilledList.last.item, apiFilledList.last.item)
}

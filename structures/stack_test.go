package structures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack(t *testing.T) {
	stack := new(Stack)
	var item StackElement

	stack.Push("рцы")
	stack.Push("слово")
	stack.Push("твёрдо")
	assert.Equal(t, "твёрдо", stack.last.item)
	assert.Equal(t, "слово", stack.last.next.item)
	assert.Equal(t, 3, stack.length)

	item = stack.Pop()
	assert.Equal(t, item, "твёрдо")
	item = stack.Pop()
	assert.Equal(t, item, "слово")
	item = stack.Pop()
	assert.Equal(t, item, "рцы")
	assert.Equal(t, 0, stack.length)
	assert.Nil(t, stack.last)
}

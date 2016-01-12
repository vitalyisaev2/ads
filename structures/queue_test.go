package structures

import (
	//"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue()

	// Enqueue three items
	items := []QueueElement{
		QueueElement("рцы"),
		QueueElement("слово"),
		QueueElement("твёрдо"),
	}

	q.Enqueue(&items[0])
	q.Enqueue(&items[1])
	q.Enqueue(&items[2])

	assert.Equal(t, 3, len(q.items))
	assert.Equal(t, *q.items[0], "рцы")
	assert.Equal(t, *q.items[1], "слово")
	assert.Equal(t, *q.items[2], "твёрдо")

	// And dequeue them
	var item *QueueElement
	item = q.Dequeue()
	assert.Equal(t, *item, "рцы")
	item = q.Dequeue()
	assert.Equal(t, *item, "слово")
	item = q.Dequeue()
	assert.Equal(t, *item, "твёрдо")
	assert.Equal(t, 0, len(q.items))
}

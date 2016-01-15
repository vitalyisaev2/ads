package structures

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue(2, 2)

	// Enqueue three items
	items := []QueueElement{
		QueueElement("рцы"),
		QueueElement("слово"),
		QueueElement("твёрдо"),
	}

	q.Enqueue(&items[0])
	q.Enqueue(&items[1])
	q.Enqueue(&items[2])

	// And dequeue them
	var item *QueueElement
	item = <-q.Dequeue()
	assert.Equal(t, *item, "рцы")
	item = <-q.Dequeue()
	assert.Equal(t, *item, "слово")
	item = <-q.Dequeue()
	assert.Equal(t, *item, "твёрдо")
}

var queueBenchmarkElement *QueueElement

func BenchmarkQueue(b *testing.B) {

	var r *QueueElement
	q := NewQueue(128, 2)

	for n := 0; n < b.N; n++ {
		ch := make(chan bool, 2)
		go func() {
			for i := 0; i < math.MaxUint16; i++ {
				item := QueueElement(i)
				q.Enqueue(&item)
			}
			ch <- true
		}()
		go func() {
			for i := 0; i < math.MaxUint16; i++ {
				r = <-q.Dequeue()
			}
			ch <- true
		}()
		<-ch
		<-ch
	}
	queueBenchmarkElement = r
}

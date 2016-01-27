package structures

import (
	//"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

const (
	maxQueueItemsAmount = math.MaxUint8
	defaultCapacity     = 1 << 10
)

//---------------------- Tests ------------------------

// Send and check Int values
func queueIntValues(q Queue, t *testing.T) {
	ch := make(chan bool, 2)
	go func() {
		for i := 0; i < maxQueueItemsAmount; i++ {
			q.Enqueue(i)
		}
		ch <- true
	}()
	go func() {
		var r int
		for i := 0; i < maxQueueItemsAmount; i++ {
			r = (<-q.Dequeue()).(int)
			assert.Equal(t, i, r)
		}
		ch <- true
	}()
	<-ch
	<-ch
	assert.Equal(t, 0, q.Len())
}

// Send and check Int pointers
func queueIntPointers(q Queue, t *testing.T) {
	ch := make(chan bool, 2)
	go func() {
		for i := 0; i < maxQueueItemsAmount; i++ {
			x := i
			q.Enqueue(&x)
		}
		ch <- true
	}()
	go func() {
		var r *int
		for i := 0; i < maxQueueItemsAmount; i++ {
			r = (<-q.Dequeue()).(*int)
			assert.Equal(t, i, *r)
		}
		ch <- true
	}()
	<-ch
	<-ch
	assert.Equal(t, 0, q.Len())
}

func TestChannelQueueIntValues(t *testing.T) {
	q := NewQueue("channelQueue", defaultCapacity)
	queueIntValues(q, t)
}

func TestChannelQueueIntPointers(t *testing.T) {
	q := NewQueue("channelQueue", defaultCapacity)
	queueIntPointers(q, t)
}

// ------------------ Benchmark  -------------------

var queueBenchmarkElement interface{}

func benchmarkQueue(q Queue, b *testing.B) {

	var r interface{}

	for n := 0; n < b.N; n++ {
		ch := make(chan bool, 2)
		go func() {
			for i := 0; i < maxQueueItemsAmount; i++ {
				q.Enqueue(i)
			}
			ch <- true
		}()
		go func() {
			for i := 0; i < maxQueueItemsAmount; i++ {
				r = <-q.Dequeue()
			}
			ch <- true
		}()
		<-ch
		<-ch
	}
	queueBenchmarkElement = r
}

func BenchmarkChannelQueue(b *testing.B) {
	q := NewQueue("channelQueue", defaultCapacity)
	benchmarkQueue(q, b)
}

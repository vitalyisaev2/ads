package structures

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

const (
	maxQueueItemsAmount = math.MaxUint16
	defaultCapacity     = 1 << 4
)

//---------------------- Tests ------------------------

// Send and check Int values/pointers concurrently
func queueConcurrentTest(t *testing.T, q Queue, regime string) {
	ch := make(chan bool, 2)
	go func() {
		for i := 0; i < maxQueueItemsAmount; i++ {
			switch regime {
			case "values":
				q.Enqueue(i)
			case "pointers":
				x := i
				q.Enqueue(&x)
			}
		}
		ch <- true
	}()
	go func() {
		for i := 0; i < maxQueueItemsAmount; i++ {
			switch regime {
			case "values":
				r := (<-q.Dequeue()).(int)
				assert.Equal(t, i, r)
			case "pointers":
				r := (<-q.Dequeue()).(*int)
				assert.Equal(t, i, *r)
			}
		}
		ch <- true
	}()
	<-ch
	<-ch
	assert.Equal(t, 0, q.Len())
}

// Store values, than flush it
func queueSequentialTest(t *testing.T, q Queue, regime string) {

	for i := 0; i < maxQueueItemsAmount; i++ {
		switch regime {
		case "values":
			q.Enqueue(i)
		case "pointers":
			x := i
			q.Enqueue(&x)
		}
	}

	for i := 0; i < maxQueueItemsAmount; i++ {
		switch regime {
		case "values":
			r := (<-q.Dequeue()).(int)
			assert.Equal(t, i, r)
		case "pointers":
			r := (<-q.Dequeue()).(*int)
			assert.Equal(t, i, *r)
		}
		//t.Log("After dequeue: ", q.Len())
	}

	assert.Equal(t, 0, q.Len())
}

func TestChannelQueue(t *testing.T) {
	var q Queue

	q = NewQueue("channelQueue", defaultCapacity)
	queueSequentialTest(t, q, "values")

	q = NewQueue("channelQueue", defaultCapacity)
	queueSequentialTest(t, q, "pointers")

	q = NewQueue("channelQueue", defaultCapacity)
	queueConcurrentTest(t, q, "values")

	q = NewQueue("channelQueue", defaultCapacity)
	queueConcurrentTest(t, q, "pointers")
}

func TestSliceQueue(t *testing.T) {
	var q Queue

	q = NewQueue("sliceQueue", defaultCapacity)
	queueSequentialTest(t, q, "values")

	q = NewQueue("sliceQueue", defaultCapacity)
	queueSequentialTest(t, q, "pointers")

	q = NewQueue("sliceQueue", defaultCapacity)
	queueConcurrentTest(t, q, "values")

	q = NewQueue("sliceQueue", defaultCapacity)
	queueConcurrentTest(t, q, "pointers")
}

func TestLinkedListQueue(t *testing.T) {
	var q Queue

	q = NewQueue("channelQueue", defaultCapacity)
	queueSequentialTest(t, q, "values")

	q = NewQueue("channelQueue", defaultCapacity)
	queueSequentialTest(t, q, "pointers")

	q = NewQueue("channelQueue", defaultCapacity)
	queueConcurrentTest(t, q, "values")

	q = NewQueue("channelQueue", defaultCapacity)
	queueConcurrentTest(t, q, "pointers")
}

// ------------------ Benchmarks -------------------

var queueBenchmarkElement interface{}

func queueConcurrentBenchmark(b *testing.B, q Queue) {

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

func queueSequentialBenchmark(b *testing.B, q Queue) {

	var r interface{}

	for n := 0; n < b.N; n++ {
		for i := 0; i < maxQueueItemsAmount; i++ {
			q.Enqueue(i)
		}
		for i := 0; i < maxQueueItemsAmount; i++ {
			r = <-q.Dequeue()
		}
	}
	queueBenchmarkElement = r
}

// Concurrent benchmarks
func BenchmarkChannelQueueConcurrent16(b *testing.B) {
	q := NewQueue("channelQueue", 16)
	queueConcurrentBenchmark(b, q)
}

func BenchmarkSliceQueueConcurrent16(b *testing.B) {
	q := NewQueue("sliceQueue", 16)
	queueConcurrentBenchmark(b, q)
}

func BenchmarkChannelQueueConcurrent256(b *testing.B) {
	q := NewQueue("channelQueue", 256)
	queueConcurrentBenchmark(b, q)
}

func BenchmarkSliceQueueConcurrent256(b *testing.B) {
	q := NewQueue("sliceQueue", 256)
	queueConcurrentBenchmark(b, q)
}

func BenchmarkChannelQueueConcurrent4096(b *testing.B) {
	q := NewQueue("channelQueue", 4096)
	queueConcurrentBenchmark(b, q)
}

func BenchmarkSliceQueueConcurrent4096(b *testing.B) {
	q := NewQueue("sliceQueue", 4096)
	queueConcurrentBenchmark(b, q)
}

func BenchmarkLinkedListQueueConcurrent0(b *testing.B) {
	q := NewQueue("linkedListQueue", 0)
	queueConcurrentBenchmark(b, q)
}

// Sequential benchmarks
func BenchmarkChannelQueueSequential16(b *testing.B) {
	q := NewQueue("channelQueue", 16)
	queueSequentialBenchmark(b, q)
}

func BenchmarkSliceQueueSequential16(b *testing.B) {
	q := NewQueue("sliceQueue", 16)
	queueSequentialBenchmark(b, q)
}

func BenchmarkChannelQueueSequential256(b *testing.B) {
	q := NewQueue("channelQueue", 256)
	queueSequentialBenchmark(b, q)
}

func BenchmarkSliceQueueSequential256(b *testing.B) {
	q := NewQueue("sliceQueue", 256)
	queueSequentialBenchmark(b, q)
}

func BenchmarkChannelQueueSequential4096(b *testing.B) {
	q := NewQueue("channelQueue", 4096)
	queueSequentialBenchmark(b, q)
}

func BenchmarkSliceQueueSequential4096(b *testing.B) {
	q := NewQueue("sliceQueue", 4096)
	queueSequentialBenchmark(b, q)
}

func BenchmarkLinkedListQueueSequentila0(b *testing.B) {
	q := NewQueue("linkedListQueue", 0)
	queueSequentialBenchmark(b, q)
}

// Skiena 3.2: stacks and queues
// Three implementations of thread-safe queues backed with a
// channel of variable capacity, slice and doubly linked list

package structures

import (
	"log"
	//"runtime"
)

// ----------------------- Queue -------------------------------

type Queue interface {
	Enqueue(interface{})
	Dequeue() <-chan interface{}
	Len() int
	poll()
}

// Multiplication coeffitient in resize operations
const multiplier = 2

// ---------------------- Basic Queue -------------------------
type commonChannel chan interface{}

// basicQueue has no Poll() implementation
type basicQueue struct {
	enqueue   commonChannel
	dequeue   chan commonChannel
	occupancy chan struct{}
}

// Pushes item to the end of the Queue
func (q *basicQueue) Enqueue(item interface{}) {
	q.enqueue <- item
}

// Pops item from the beginning of the Queue
func (q *basicQueue) Dequeue() <-chan interface{} {
	ch := make(chan interface{})
	q.dequeue <- ch
	return ch
}

// -------------- Queue with channel-based synchronization and commonChannel as a buffer -------------------------

type channelQueue struct {
	basicQueue
	queue    commonChannel
	lock     chan struct{}
	capacity int
}

func (q *channelQueue) poll() {

	// First unlock the queue for then enqueuer (like mutex opening)
	q.lock <- struct{}{}

	// Enqueuer
	go func() {
		defer func() {
			close(q.enqueue)
			close(q.queue)
		}()
		for {
			item := <-q.enqueue

			// Lock channel
			<-q.lock

			// Resize channel if there's no enough space
			if len(q.queue) == cap(q.queue) {
				q.bufferIncrease()
			}

			// Enqueue item
			q.queue <- item

			// Release lock
			q.lock <- struct{}{}

			// Tell to the Dequeuer that there's something to dequeue
			go func() {
				q.occupancy <- struct{}{}
			}()

		}
	}()

	// Dequeuer (will block if there's nothing to read)
	go func() {
		defer func() {
			close(q.dequeue)
		}()
		for {
			ch := <-q.dequeue

			// Wait till channel will contain some data
			<-q.occupancy

			// Wait for mutex
			<-q.lock

			// Read element and send it to user
			item := <-q.queue
			ch <- item
			close(ch)

			// Decrease main channel capacity in order to free unused memory
			// (but not so much)
			queueCap, queueLen := cap(q.queue), len(q.queue)
			queueCapDecreased := int(queueCap / multiplier)
			if queueLen < queueCapDecreased && queueCapDecreased >= q.capacity {
				q.bufferDecrease()
			}

			q.lock <- struct{}{}
		}
	}()
}

// Replace the channel with another one with capacity increased in
// multiplier times
func (q *channelQueue) bufferIncrease() {
	substituteCommonChannel := make(commonChannel, cap(q.queue)*multiplier)
	items := len(q.queue)
	for i := 0; i < items; i++ {
		item := <-q.queue
		substituteCommonChannel <- item
	}
	q.queue = substituteCommonChannel
}

// Replace the channel with another one with capacity decreased in
// divider times
func (q *channelQueue) bufferDecrease() {
	substituteCommonChannel := make(commonChannel, int(cap(q.queue)/multiplier))
	items := len(q.queue)
	for i := 0; i < items; i++ {
		item := <-q.queue
		substituteCommonChannel <- item
	}
	q.queue = substituteCommonChannel
}

// Returns the length of inner channel
func (q *channelQueue) Len() int {
	<-q.lock
	defer func() {
		q.lock <- struct{}{}
	}()
	return len(q.queue)
}

// -------------- Queue with channel-based synchronization and []interface{} as a buffer -------------------------
type sliceQueue struct {
	basicQueue
	queue    []interface{}
	lock     chan struct{}
	capacity int
}

func (q *sliceQueue) poll() {
	// First unlock the queue for then enqueuer (like mutex opening)
	q.lock <- struct{}{}

	// Enqueuer
	go func() {
		defer func() {
			close(q.enqueue)
		}()
		for {
			item := <-q.enqueue

			// Lock queue
			<-q.lock
			// Send new item to slice. Golang will change the capacity if needed
			q.queue = append(q.queue, item)
			// Unlock queue
			q.lock <- struct{}{}

			// Tell to the Dequeuer that there's something to dequeue
			go func() {
				q.occupancy <- struct{}{}
			}()
		}
	}()

	// Dequeuer
	go func() {
		defer func() {
			close(q.dequeue)
		}()
		for {
			ch := <-q.dequeue

			// Wait till channel will contain some data
			<-q.occupancy

			// Unlock queue
			<-q.lock

			// Read element and send it to user
			item := q.queue[0]
			ch <- item
			close(ch)

			// Remove element from slice
			q.queue = append(q.queue[:0], q.queue[1:]...)

			// Free memory if needed
			queueCap, queueLen := cap(q.queue), len(q.queue)
			queueCapDecreased := int(queueCap / multiplier)
			if queueLen < queueCapDecreased && queueCapDecreased >= q.capacity {
				queue := make([]interface{}, queueLen, queueCapDecreased)
				copy(queue, q.queue[:queueLen])
				q.queue = queue
			}
			q.lock <- struct{}{}
		}
	}()
}

// Returns the length of inner channel
func (q *sliceQueue) Len() int {
	<-q.lock
	defer func() {
		q.lock <- struct{}{}
	}()
	return len(q.queue)
}

// --------------- Queue with no initial capacity backed with a (hopefully) thread-safe DoublyLinkedList ------------------
type linkedListQueue struct {
	basicQueue
	queue DoublyLinkedList
}

func (q *linkedListQueue) poll() {

	// Enqueuer
	go func() {
		defer func() {
			close(q.enqueue)
		}()
		for {
			item := <-q.enqueue
			err := q.queue.PushBack(item)
			if err != nil {
				log.Panicln("Error in DoublyLinkedList.PushBack(): ", err)
			}

			go func() {
				q.occupancy <- struct{}{}
			}()
		}
	}()

	// Dequeuer
	go func() {
		defer func() {
			close(q.dequeue)
		}()
		for {
			ch := <-q.dequeue
			<-q.occupancy

			item, err := q.queue.PopFront()
			if err != nil {
				log.Panicln("Error in DoublyLinkedList.PopFront(): ", err)
			}

			ch <- item
			close(ch)
		}
	}()
}

// Returns the length of inner channel
func (q *linkedListQueue) Len() int {
	return q.queue.Len()
}

// ---------------------------- Queue fabric ----------------------------------------
func NewQueue(kind string, capacity int) Queue {
	enqueue := make(commonChannel)
	dequeue := make(chan commonChannel)
	occupancy := make(chan struct{})
	lock := make(chan struct{}, 1)

	var q Queue
	if kind == "channelQueue" {
		queue := make(commonChannel, capacity)
		q = &channelQueue{
			basicQueue{enqueue, dequeue, occupancy},
			queue, lock, capacity,
		}
	} else if kind == "sliceQueue" {
		queue := make([]interface{}, 0, capacity)
		q = &sliceQueue{
			basicQueue{enqueue, dequeue, occupancy},
			queue, lock, capacity,
		}
	} else if kind == "linkedListQueue" {
		queue := NewDoublyLinkedList()
		q = &linkedListQueue{
			basicQueue{enqueue, dequeue, occupancy},
			queue,
		}
	}
	q.poll()
	return q
}

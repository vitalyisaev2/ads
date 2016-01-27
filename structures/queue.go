// Skiena 3.2: stacks and queues

package structures

import (
//"sync"
//"fmt"
)

// ----------------------- Queue -------------------------------

type Queue interface {
	Enqueue(interface{})
	Dequeue() <-chan interface{}
	Len() int
	poll()
}

// ---------------------- ResizableChannel  -------------------------

// ResizableChannel is a channel of variable capacity that underlies
// channelQueue
type ResizableChannel chan interface{}

// Multiplication coeffitient in resize operations
const multiplier = 2

// Replace the channel with another one with capacity increased in
// multiplier times
func (qc ResizableChannel) bufferIncrease(multiplier int) {
	capacity := cap(qc)
	new_qc := make(chan interface{}, capacity*multiplier)
	for i := 0; i < len(new_qc); i++ {
		item := <-qc
		new_qc <- item
	}
	qc = new_qc
}

// Replace the channel with another one with capacity decreased in
// divider times
func (qc ResizableChannel) bufferDecrease(divider int) {
	capacity := cap(qc)
	new_qc := make(chan interface{}, int(capacity/divider))
	for i := 0; i < len(new_qc); i++ {
		item := <-qc
		new_qc <- item
	}
	qc = new_qc
}

// ---------------------- Basic Queue -------------------------

// basicQueue has no Poll() implementation
type basicQueue struct {
	enqueue   ResizableChannel
	dequeue   chan ResizableChannel
	occupancy chan struct{}
	lock      chan struct{}
	capacity  int
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

// -------------- Queue with channel-based synchronization with ResizableChannel as a buffer -------------------------

type channelQueue struct {
	basicQueue
	queue ResizableChannel
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

			// Resize channel if there's no enough space
			if len(q.queue) == cap(q.queue) {
				<-q.lock
				q.queue.bufferIncrease(multiplier)
				q.lock <- struct{}{}
			}

			// Enqueue item
			q.queue <- item

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

			// Read element and send it to user
			item := <-q.queue
			ch <- item
			close(ch)

			// Decrease main channel capacity in order to free unused memory
			queueCap, queueLen := cap(q.queue), len(q.queue)
			if queueLen < int(queueCap/multiplier) && (int(queueCap/multiplier) > q.capacity) {
				<-q.lock
				q.queue.bufferDecrease(multiplier)
				q.lock <- struct{}{}
			}
		}
	}()
}

// Returns the length of inner channel
func (q *channelQueue) Len() int {
	<-q.lock
	defer func() {
		q.lock <- struct{}{}
	}()
	return len(q.queue)
}

// --------------------------------
type sliceQueue struct {
	basicQueue
	queue []interface{}
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
				queue := make([]interface{}, queueCapDecreased)
				copy(queue, q.queue[:queueCap])
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

// ---------------------------------- Queue fabric -------------------------

// Queue fabric function
func NewQueue(kind string, capacity int) Queue {
	enqueue := make(ResizableChannel)
	dequeue := make(chan ResizableChannel)
	occupancy := make(chan struct{})
	lock := make(chan struct{}, 1)

	var q Queue
	if kind == "channelQueue" {
		queue := make(ResizableChannel, capacity)
		q = &channelQueue{
			basicQueue{enqueue, dequeue, occupancy, lock, capacity},
			queue,
		}
	} else if kind == "sliceQueue" {
		queue := make([]interface{}, 0, capacity)
		q = &sliceQueue{
			basicQueue{enqueue, dequeue, occupancy, lock, capacity},
			queue,
		}
	}
	q.poll()
	return q
}

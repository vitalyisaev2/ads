// Skiena 3.2: stacks and queues

package structures

import (
//"sync"
)

// ----------------------- Queue -------------------------------
const multiplier = 2

type Queue interface {
	Enqueue(interface{})
	Dequeue() <-chan interface{}
	Len() int
	poll()
}

// ---------------------- Queue Channel -------------------------

// ResizableChannel is a resizable channel of interface{} type
type ResizableChannel chan interface{}

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
	enqueue  ResizableChannel
	dequeue  chan ResizableChannel
	capacity int
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

// -------------- Queue with channel-based synchronization (based on ResizableChannel) -------------------------

type channelQueue struct {
	basicQueue
	queue     ResizableChannel
	occupancy chan struct{}
	lock      chan struct{}
}

func (q *channelQueue) poll() {

	// First unlock the queue for then enqueuer
	q.lock <- struct{}{}

	// Enqueuer
	go func() {
		defer func() {
			close(q.enqueue)
			close(q.queue)
		}()
		for {
			select {

			case item := <-q.enqueue:

				// Resize channel if there's no enough space
				if len(q.queue) == cap(q.queue) {
					<-q.lock
					q.queue.bufferIncrease(multiplier)
					q.lock <- struct{}{}
				}

				// Enqueue item
				q.queue <- item

				// Tell to the Dequeuer that there's something in channel
				go func() {
					q.occupancy <- struct{}{}
				}()

			}
		}
	}()

	// Dequeuer (will block if there's nothing to read)
	go func() {
		defer func() {
			close(q.dequeue)
		}()
		for {
			select {
			case ch := <-q.dequeue:

				// Wait till channel will contain some data
				<-q.occupancy

				// Read element and send it to user
				item := <-q.queue
				ch <- item
				close(ch)

				// Decrease main channel capacity in order to free unused memory
				queueCap, queueLen := cap(q.queue), len(q.queue)
				if queueLen < int(queueCap/multiplier) && (queueCap*multiplier > q.capacity) {
					<-q.lock
					q.queue.bufferDecrease(multiplier)
					q.lock <- struct{}{}
				}
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

// ---------------------------------- Queue fabric -------------------------

// Queue fabric function
func NewQueue(kind string, capacity int) Queue {
	enqueue := make(ResizableChannel)
	dequeue := make(chan ResizableChannel)

	var q Queue
	if kind == "channelQueue" {
		queue := make(ResizableChannel, capacity)
		occupancy := make(chan struct{})
		lock := make(chan struct{}, 1)

		q = &channelQueue{
			basicQueue{enqueue, dequeue, capacity},
			queue,
			occupancy,
			lock,
		}

		q.poll()
	}
	return q
}

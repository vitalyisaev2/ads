// Skiena 3.2: stacks and queues

package structures

import (
	"sync"
)

// QueueChannel is a resizable channel of interface{} type
type QueueChannel chan interface{}

// Replace the channel with another one with capacity increased in
// multiplier times
func (qc QueueChannel) bufferIncrease(multiplier int) {
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
func (qc QueueChannel) bufferDecrease(divider int) {
	capacity := cap(qc)
	new_qc := make(chan interface{}, int(capacity/divider))
	for i := 0; i < len(new_qc); i++ {
		item := <-qc
		new_qc <- item
	}
	qc = new_qc
}

// Queue interface
type Queue interface {
	Poll()
	Enqueue(interface{})
	Dequeue() <-chan interface{}
	Len() int
}

// basicQueue has no Poll() implementation
type basicQueue struct {
	enqueue    QueueChannel
	queue      QueueChannel
	dequeue    chan QueueChannel
	occupancy  chan struct{}
	capacity   int
	multiplier int
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

func (q *basicQueue) Len() int {
	return len(q.queue)
}

// In a mutexQueue mutex is used for preserving main channel
// from the corruption
type mutexQueue struct {
	basicQueue
	mutex sync.Mutex
}

func (q *mutexQueue) Poll() {

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
					q.mutex.Lock()
					q.queue.bufferIncrease(q.multiplier)
					q.mutex.Unlock()
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
				if queueLen < int(queueCap/q.multiplier) && (queueCap*q.multiplier > q.capacity) {
					q.mutex.Lock()
					q.queue.bufferDecrease(q.multiplier)
					q.mutex.Unlock()
				}
			}
		}
	}()
}

func NewQueue(kind string, capacity, multiplier int) Queue {
	enqueue := make(QueueChannel)
	queue := make(QueueChannel, capacity)
	dequeue := make(chan QueueChannel)
	occupancy := make(chan struct{})
	mutex := sync.Mutex{}

	var q Queue
	if kind == "mutex" {
		q = &mutexQueue{
			basicQueue{enqueue, queue, dequeue, occupancy, capacity, multiplier},
			mutex,
		}
	}
	q.Poll()
	return q
}

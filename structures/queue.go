// Skiena 3.2: stacks and queues

package structures

import (
	//"fmt"
	"sync"
)

type QueueChannel chan interface{}

func (qc QueueChannel) bufferIncrease(multiplier int) {
	capacity := cap(qc)
	new_qc := make(chan interface{}, capacity*multiplier)
	for i := 0; i < len(new_qc); i++ {
		item := <-qc
		new_qc <- item
	}
	qc = new_qc
}

func (qc QueueChannel) bufferDecrease(divider int) {
	capacity := cap(qc)
	new_qc := make(chan interface{}, int(capacity/divider))
	for i := 0; i < len(new_qc); i++ {
		item := <-qc
		new_qc <- item
	}
	qc = new_qc
}

type mutexQueue struct {
	enqueue    QueueChannel
	queue      QueueChannel
	dequeue    chan QueueChannel
	occupancy  chan struct{}
	capacity   int
	multiplier int
	mutex      sync.Mutex
}

func NewMutexQueue(capacity, multiplier int) *mutexQueue {
	enqueue := make(QueueChannel)
	queue := make(QueueChannel, capacity)
	dequeue := make(chan QueueChannel)
	occupancy := make(chan struct{})
	mutex := sync.Mutex{}

	q := &mutexQueue{
		enqueue, queue, dequeue, occupancy, capacity, multiplier, mutex,
	}
	q.Poll()
	return q
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

func (q *mutexQueue) Enqueue(item interface{}) {
	q.enqueue <- item
}

func (q *mutexQueue) Dequeue() <-chan interface{} {
	ch := make(chan interface{})
	q.dequeue <- ch
	return ch
}

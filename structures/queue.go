// Skiena 3.2: stacks and queues

package structures

import (
	"sync"
)

///////////// Resizable channel is a part of queue ////////////////

type QueueElement interface{}

type QueueChannel chan *QueueElement

func (qc QueueChannel) bufferIncrease(multiplier int) {
	capacity := cap(qc)
	new_qc := make(QueueChannel, capacity*multiplier)
	for i := 0; i < len(new_qc); i++ {
		item := <-qc
		new_qc <- item
	}
	qc = new_qc
}

func (qc QueueChannel) bufferDecrease(multiplier int) {
	capacity := cap(qc)
	new_qc := make(QueueChannel, int(capacity/multiplier))
	for i := 0; i < len(new_qc); i++ {
		item := <-qc
		new_qc <- item
	}
	qc = new_qc
}

type Queue struct {
	enqueue    QueueChannel
	queue      QueueChannel
	dequeue    chan QueueChannel
	occupancy  chan bool
	capacity   int
	multiplier int
	mutex      sync.Mutex
}

func NewQueue(capacity, multiplier int) *Queue {
	enqueue := make(QueueChannel)
	queue := make(QueueChannel, capacity)
	dequeue := make(chan QueueChannel)
	occupancy := make(chan bool)
	mutex := sync.Mutex{}

	q := &Queue{
		enqueue, queue, dequeue, occupancy, capacity, multiplier, mutex,
	}
	q.Poll()
	return q
}

func (q Queue) Poll() {

	const t = true

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
					q.occupancy <- t
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

				//if len(q.queue) < int(cap(q.queue)/q.multiplier) && (cap(q.queue)*q.multiplier > q.capacity) {
				if len(q.queue) < int(cap(q.queue)/q.multiplier) {
					q.mutex.Lock()
					q.queue.bufferDecrease(q.multiplier)
					q.mutex.Unlock()
				}
			}
		}
	}()
}

func (q *Queue) Enqueue(item *QueueElement) {
	q.enqueue <- item
}

func (q *Queue) Dequeue() <-chan *QueueElement {
	ch := make(QueueChannel)
	q.dequeue <- ch
	return ch
}

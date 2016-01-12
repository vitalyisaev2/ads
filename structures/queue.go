// Skiena 3.2: stacks and queues

package structures

import (
//"fmt"
)

type QueueElement interface{}

type Queue struct {
	items           []*QueueElement
	enqueueRequests chan *QueueElement
	dequeueRequests chan interface{}
	dequeueResults  chan *QueueElement
}

func (q *Queue) Enqueue(item *QueueElement) {
	q.enqueueRequests <- item
}

func (q *Queue) Dequeue() *QueueElement {
	var request struct{}
	q.dequeueRequests <- request
	return <-q.dequeueResults
}

func (q *Queue) poll() {
	defer func() {
		close(q.enqueueRequests)
		close(q.dequeueResults)
		close(q.dequeueResults)
	}()

	for {

		select {
		case item := <-q.enqueueRequests:
			q.items = append(q.items, item)
		case <-q.dequeueRequests:
			q.dequeueResults <- q.items[0]
			q.items = q.items[1:]
		}
	}
}

func NewQueue() *Queue {
	q := new(Queue)
	q.enqueueRequests = make(chan *QueueElement)
	q.dequeueRequests = make(chan interface{})
	q.dequeueResults = make(chan *QueueElement)

	go q.poll()
	return q
}

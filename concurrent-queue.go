package main

import "sync"

// ConcurrentQueue - thread safe queue
type ConcurrentQueue struct {
	sync.Mutex
	Items []interface{}
}

// NewQueue - constructor
func NewQueue(capacity int) *ConcurrentQueue {
	return &ConcurrentQueue{
		Items: make([]interface{}, 0, capacity),
	}
}

// Push - push item to the concurrent queue
func (q *ConcurrentQueue) Push(item interface{}) {
	q.Lock()
	defer q.Unlock()
	q.Items = append(q.Items, item)
}

// Pop - pop item from concurrent queue
func (q *ConcurrentQueue) Pop() interface{} {
	q.Lock()
	defer q.Unlock()
	if len(q.Items) == 0 {
		return nil
	}
	item := q.Items[0]
	q.Items[0] = nil
	q.Items = q.Items[1:]
	return item
}

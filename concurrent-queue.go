package main

import "sync"

type concurrentQueue struct {
	sync.Mutex
	Items []*initialCondition
}

func newQueue(capacity int) *concurrentQueue {
	return &concurrentQueue{
		Items: make([]*initialCondition, 0, capacity),
	}
}

func (q *concurrentQueue) push(item *initialCondition) {
	q.Lock()
	defer q.Unlock()
	q.Items = append(q.Items, item)
}

func (q *concurrentQueue) pop() *initialCondition {
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

package main

// Queue interface.
type Queue interface {
	Push(key interface{})
	Pop() interface{}
	Contains(key interface{}) bool
	Len() int
	Keys() []interface{}
}

// SimpleQueue defines Queue implementations.
type SimpleQueue struct {
	size  int
	queue []interface{}
}

// New creates queue with given size.
func New(size int) Queue {
	return &SimpleQueue{
		size: size,
	}
}

// Push new key into queue.
func (q *SimpleQueue) Push(key interface{}) {
	// log.Println(len(q.queue), q.size)
	if len(q.queue) == q.size {
		q.Pop()
	}

	q.queue = append(q.queue, key)
}

// Pop new key into queue.
func (q *SimpleQueue) Pop() interface{} {
	removed := q.queue[0]
	q.queue = q.queue[1:len(q.queue)]
	return removed
}

// Contains .
func (q *SimpleQueue) Contains(key interface{}) bool {
	for _, val := range q.queue {
		if val == key {
			return true
		}
	}

	return false
}

// Len .
func (q *SimpleQueue) Len() int {
	return len(q.queue)
}

// Keys .
func (q *SimpleQueue) Keys() []interface{} {
	return q.queue
}

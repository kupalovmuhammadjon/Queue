package queue

import "sync"

// Queue interface defines the methods that any queue implementation should have
type Queue interface {
	Push(val interface{})
	Pop() interface{}
	Peek() interface{}
	IsEmpty() bool
	Size() int
}

// node represents a node in the queue
type node struct {
	val  interface{}
	next *node
}

type linkedQueue struct {
    mu    sync.Mutex
	front *node
	rear  *node
	len   int
}

// NewQueue creates and returns a new Queue
func NewQueue() Queue {
    return &linkedQueue{}
}

//
func (q *linkedQueue) Push(val interface{}) {
	q.mu.Lock()
	defer q.mu.Unlock()
	newNode := &node{val: val}

	if q.IsEmpty() {
		q.front = newNode
		q.rear = newNode
	} else {
		q.rear.next = newNode
		q.rear = newNode
	}

	q.len++
}

func (q *linkedQueue) Pop() interface{} {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.IsEmpty() {
		return nil
	}

	val := q.front.val
	q.front = q.front.next

	q.len--

	if q.IsEmpty() {
		q.rear = nil
	}

	return val
}

func (q *linkedQueue) IsEmpty() bool {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.len == 0
}

func (q *linkedQueue) Peek() interface{} {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.IsEmpty() {
		return nil
	}
	return q.front.val
}

func (q *linkedQueue) Size() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.len
}

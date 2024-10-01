package queue

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

// linkedQueue is an implementation of Queue using a linked list
type linkedQueue struct {
    front *node
    rear  *node
    len   int
}

// NewQueue creates and returns a new Queue
func NewQueue() Queue {
    return &linkedQueue{}
}

// Push adds a new element to the rear of the Queue
func (q *linkedQueue) Push(val interface{}) {
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

// Pop removes and returns the front element of the Queue
func (q *linkedQueue) Pop() interface{} {
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

// IsEmpty returns true if the Queue is empty
func (q *linkedQueue) IsEmpty() bool {
    return q.len == 0
}

// Peek returns the front element of the Queue without removing it
func (q *linkedQueue) Peek() interface{} {
    if q.IsEmpty() {
        return nil
    }
    return q.front.val
}

// Size returns the number of elements in the Queue
func (q *linkedQueue) Size() int {
    return q.len
}
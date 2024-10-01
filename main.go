package main

type Node struct {
	Val  interface{}
	Next *Node
}

type Queue struct {
	Front *Node
	Rear  *Node
	Len   int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Push(val interface{}) {
	newNode := &Node{Val: val}

	if q.IsEmpty() {
		q.Front = newNode
		q.Rear = newNode
	} else {
		q.Rear.Next = newNode
		q.Rear = newNode
	}

	q.Len++
}

func (q *Queue) Pop() interface{} {
	if q.IsEmpty() {
		return nil
	}

	val := q.Front.Val
	q.Front = q.Front.Next

	q.Len--

	if q.IsEmpty() {
		q.Rear = nil
	}

	return val
}

func (q *Queue) IsEmpty() bool {
	return q.Len == 0
}

func (q *Queue) Peek() interface{} {
	if q.IsEmpty() {
		return nil
	}
	return q.Front.Val
}

func (q *Queue) Size() int {
	return q.Len
}

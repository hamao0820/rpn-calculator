package ds

import "fmt"

type Queue[T any] struct {
	data []T
}

// NewQueue instantiates a new queue
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{data: make([]T, 0)}
}

// Push adds a new element at the end of the queue
func (q *Queue[T]) Push(n T) {
	q.data = append(q.data, n)
}

// Pop removes the first element from queue
func (q *Queue[T]) Pop() bool {
	if q.IsEmpty() {
		return false
	}
	var defaultVal T
	q.data[0] = defaultVal
	q.data = q.data[1:]
	return true
}

// Front returns the first element of queue
func (q *Queue[T]) Front() T {
	return q.data[0]
}

// IsEmpty checks if the queue is empty
func (q *Queue[T]) IsEmpty() bool {
	return len(q.data) == 0
}

// String implements Stringer interface
func (q *Queue[T]) String() string {
	return fmt.Sprint(q.data)
}

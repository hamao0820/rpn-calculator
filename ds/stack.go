package ds

import "fmt"

type Stack[T any] struct {
	data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		data: make([]T, 0),
	}
}

func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
}

func (s *Stack[T]) Pop() bool {
	if s.IsEmpty() {
		return false
	}
	s.data = s.data[:len(s.data)-1]
	return true
}

func (s *Stack[T]) Top() T {
	return s.data[len(s.data)-1]
}

func (s *Stack[T]) Len() int {
	return len(s.data)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack[T]) String() string {
	return fmt.Sprint(s.data)
}

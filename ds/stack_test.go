package ds

import (
	"testing"
)

func TestStackInt(t *testing.T) {
	stack := NewStack[int]()

	if !stack.IsEmpty() {
		t.Fatalf("True is expected, but %v\n", stack.IsEmpty())
	}

	stack.Push(10)
	stack.Push(1)
	stack.Push(-5)

	if stack.IsEmpty() {
		t.Fatalf("False is expected, but %v\n", stack.IsEmpty())
	}

	if stack.Top() != -5 {
		t.Fatalf("-5 is expected, but %v\n", stack.Top())
	}

	stack.Pop()
	if stack.Top() != 1 {
		t.Fatalf("1 is expected, but %v\n", stack.Top())
	}

	stack.Pop()
	if stack.Top() != 10 {
		t.Fatalf("10 is expected, but %v\n", stack.Top())
	}

	stack.Pop()
	if !stack.IsEmpty() {
		t.Fatalf("True is expected, but %v\n", stack.IsEmpty())
	}
}

func TestStackString(t *testing.T) {
	stack := NewStack[string]()

	if !stack.IsEmpty() {
		t.Fatalf("True is expected, but %v\n", stack.IsEmpty())
	}

	stack.Push("hello")
	stack.Push("world")
	stack.Push("golang")

	if stack.IsEmpty() {
		t.Fatalf("False is expected, but %v\n", stack.IsEmpty())
	}

	if stack.Top() != "golang" {
		t.Fatalf("golang is expected, but %v\n", stack.Top())
	}

	stack.Pop()
	if stack.Top() != "world" {
		t.Fatalf("world is expected, but %v\n", stack.Top())
	}

	stack.Pop()
	if stack.Top() != "hello" {
		t.Fatalf("hello is expected, but %v\n", stack.Top())
	}

	stack.Pop()
	if !stack.IsEmpty() {
		t.Fatalf("True is expected, but %v\n", stack.IsEmpty())
	}
}

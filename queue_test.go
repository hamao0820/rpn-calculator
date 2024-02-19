package main

import (
	"testing"
)

func TestQueueInt(t *testing.T) {
	queue := NewQueue[int]()

	if !queue.IsEmpty() {
		t.Fatalf("True is expected, but %v\n", queue.IsEmpty())
	}

	queue.Push(10)
	queue.Push(1)
	queue.Push(-5)

	if queue.IsEmpty() {
		t.Fatalf("False is expected, but %v\n", queue.IsEmpty())
	}

	if queue.Front() != 10 {
		t.Fatalf("10 is expected, but %v\n", queue.Front())
	}

	queue.Pop()
	if queue.Front() != 1 {
		t.Fatalf("1 is expected, but %v\n", queue.Front())
	}

	queue.Pop()
	if queue.Front() != -5 {
		t.Fatalf("-5 is expected, but %v\n", queue.Front())
	}

	queue.Pop()
	if !queue.IsEmpty() {
		t.Fatalf("True is expected, but %v\n", queue.IsEmpty())
	}
}

func TestQueueString(t *testing.T) {
	queue := NewQueue[string]()

	if !queue.IsEmpty() {
		t.Fatalf("True is expected, but %v\n", queue.IsEmpty())
	}

	queue.Push("hello")
	queue.Push("world")
	queue.Push("golang")

	if queue.IsEmpty() {
		t.Fatalf("False is expected, but %v\n", queue.IsEmpty())
	}

	if queue.Front() != "hello" {
		t.Fatalf("hello is expected, but %v\n", queue.Front())
	}

	queue.Pop()
	if queue.Front() != "world" {
		t.Fatalf("world is expected, but %v\n", queue.Front())
	}

	queue.Pop()
	if queue.Front() != "golang" {
		t.Fatalf("golang is expected, but %v\n", queue.Front())
	}

	queue.Pop()
	if !queue.IsEmpty() {
		t.Fatalf("True is expected, but %v\n", queue.IsEmpty())
	}
}

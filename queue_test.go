package main

import (
	"fmt"
	"testing"
)

func TestQueueInt(t *testing.T) {
	queue := NewQueue[int]()
	fmt.Println(queue)

	if !queue.IsEmpty() {
		t.Fatalf("True is expected, but %v\n", queue.IsEmpty())
	}

	queue.Push(10)
	fmt.Println(queue)
	queue.Push(1)
	fmt.Println(queue)
	queue.Push(-5)
	fmt.Println(queue)

	if queue.IsEmpty() {
		t.Fatalf("False is expected, but %v\n", queue.IsEmpty())
	}

	if queue.Front() != 10 {
		t.Fatalf("10 is expected, but %v\n", queue.Front())
	}

	queue.Pop()
	fmt.Println(queue)
	if queue.Front() != 1 {
		t.Fatalf("1 is expected, but %v\n", queue.Front())
	}

	queue.Pop()
	fmt.Println(queue)
	if queue.Front() != -5 {
		t.Fatalf("-5 is expected, but %v\n", queue.Front())
	}

	queue.Pop()
	fmt.Println(queue)
	if !queue.IsEmpty() {
		t.Fatalf("True is expected, but %v\n", queue.IsEmpty())
	}
}

func TestQueueString(t *testing.T) {
	queue := NewQueue[string]()
	fmt.Println(queue)

	if !queue.IsEmpty() {
		t.Fatalf("True is expected, but %v\n", queue.IsEmpty())
	}

	queue.Push("hello")
	fmt.Println(queue)
	queue.Push("world")
	fmt.Println(queue)
	queue.Push("golang")
	fmt.Println(queue)

	if queue.IsEmpty() {
		t.Fatalf("False is expected, but %v\n", queue.IsEmpty())
	}

	if queue.Front() != "hello" {
		t.Fatalf("hello is expected, but %v\n", queue.Front())
	}

	queue.Pop()
	fmt.Println(queue)
	if queue.Front() != "world" {
		t.Fatalf("world is expected, but %v\n", queue.Front())
	}

	queue.Pop()
	fmt.Println(queue)
	if queue.Front() != "golang" {
		t.Fatalf("golang is expected, but %v\n", queue.Front())
	}

	queue.Pop()
	fmt.Println(queue)
	if !queue.IsEmpty() {
		t.Fatalf("True is expected, but %v\n", queue.IsEmpty())
	}
}

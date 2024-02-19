package main

import (
	"fmt"
	"testing"
)

func TestStackInt(t *testing.T) {
	stack := NewStack[int]()

	fmt.Println(stack)
	if !stack.IsEmpty() {
		t.Fatalf("True is expected, but %v\n", stack.IsEmpty())
	}

	stack.Push(10)
	fmt.Println(stack)
	stack.Push(1)
	fmt.Println(stack)
	stack.Push(-5)
	fmt.Println(stack)

	if stack.IsEmpty() {
		t.Fatalf("False is expected, but %v\n", stack.IsEmpty())
	}

	if stack.Top() != -5 {
		t.Fatalf("-5 is expected, but %v\n", stack.Top())
	}

	stack.Pop()
	fmt.Println(stack)
	if stack.Top() != 1 {
		t.Fatalf("1 is expected, but %v\n", stack.Top())
	}

	stack.Pop()
	fmt.Println(stack)
	if stack.Top() != 10 {
		t.Fatalf("10 is expected, but %v\n", stack.Top())
	}

	stack.Pop()
	fmt.Println(stack)
	if !stack.IsEmpty() {
		t.Fatalf("True is expected, but %v\n", stack.IsEmpty())
	}
}

func TestStackString(t *testing.T) {
	stack := NewStack[string]()

	fmt.Println(stack)
	if !stack.IsEmpty() {
		t.Fatalf("True is expected, but %v\n", stack.IsEmpty())
	}

	stack.Push("hello")
	fmt.Println(stack)
	stack.Push("world")
	fmt.Println(stack)
	stack.Push("golang")
	fmt.Println(stack)

	if stack.IsEmpty() {
		t.Fatalf("False is expected, but %v\n", stack.IsEmpty())
	}

	if stack.Top() != "golang" {
		t.Fatalf("golang is expected, but %v\n", stack.Top())
	}

	stack.Pop()
	fmt.Println(stack)
	if stack.Top() != "world" {
		t.Fatalf("world is expected, but %v\n", stack.Top())
	}

	stack.Pop()
	fmt.Println(stack)
	if stack.Top() != "hello" {
		t.Fatalf("hello is expected, but %v\n", stack.Top())
	}

	stack.Pop()
	fmt.Println(stack)
	if !stack.IsEmpty() {
		t.Fatalf("True is expected, but %v\n", stack.IsEmpty())
	}
}

package main

import (
	"strconv"
	"strings"
)

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}

func pow(x, n int) int {
	res := 1
	for n > 0 {
		if n&1 == 1 {
			res = res * x
		}
		x = x * x
		n >>= 1
	}
	return res
}

func Calc(rpn string) int {
	stack := NewStack[int]()
	tokens := strings.Split(rpn, " ")
	for _, token := range tokens {
		switch token {
		case "+":
			{
				a := stack.Top()
				if !stack.Pop() {
					panic("stack is empty")
				}
				b := stack.Top()
				if !stack.Pop() {
					panic("stack is empty")
				}
				stack.Push(add(a, b))
			}
		case "-":
			{
				a := stack.Top()
				if !stack.Pop() {
					panic("stack is empty")
				}
				b := stack.Top()
				if !stack.Pop() {
					panic("stack is empty")
				}
				stack.Push(sub(b, a)) // 順番に注意
			}
		case "*":
			{
				a := stack.Top()
				if !stack.Pop() {
					panic("stack is empty")
				}
				b := stack.Top()
				if !stack.Pop() {
					panic("stack is empty")
				}
				stack.Push(mul(a, b))
			}
		case "/":
			{
				a := stack.Top()
				if !stack.Pop() {
					panic("stack is empty")
				}
				b := stack.Top()
				if !stack.Pop() {
					panic("stack is empty")
				}
				stack.Push(div(b, a)) // 順番に注意
			}
		case "^":
			{
				a := stack.Top()
				if !stack.Pop() {
					panic("stack is empty")
				}
				b := stack.Top()
				if !stack.Pop() {
					panic("stack is empty")
				}
				stack.Push(pow(b, a)) // 順番に注意
			}
		default:
			{
				i, err := strconv.Atoi(token)
				if err != nil {
					panic(err)
				}
				stack.Push(i)
			}
		}
	}
	return stack.Top()
}

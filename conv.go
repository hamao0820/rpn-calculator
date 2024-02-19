package main

import "strings"

var precedence = map[string]int{
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
	"^": 3,
}

func stringToTokens(s string) []string {
	tokens := make([]string, 0)
	for _, r := range strings.Split(s, " ") {
		if r == " " || r == "" {
			continue
		}
		tokens = append(tokens, string(r))
	}
	return tokens
}

func ConvRPN(s string) string {
	tokens := stringToTokens(s)
	stack := NewStack[string]()
	que := NewQueue[string]()

	for _, token := range tokens {
		if token == "+" || token == "-" || token == "*" || token == "/" || token == "^" {
			for !stack.IsEmpty() && precedence[stack.Top()] >= precedence[token] {
				que.Push(stack.Top())
				stack.Pop()
			}
			stack.Push(token)
		} else if token == "(" {
			stack.Push(token)
		} else if token == ")" {
			for stack.Top() != "(" {
				que.Push(stack.Top())
				stack.Pop()
			}
			stack.Pop()
		} else {
			que.Push(token)
		}
	}

	for !stack.IsEmpty() {
		que.Push(stack.Top())
		stack.Pop()
	}

	rpn := ""
	for !que.IsEmpty() {
		rpn += que.Front()
		rpn += " "
		que.Pop()
	}
	rpn = strings.TrimSpace(rpn)
	return rpn
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for {
		fmt.Print("input formula > ")
		// python input() equivalent in Go
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		// remove newline character
		input = input[:len(input)-1]
		fmt.Println("You entered: ", input)
		rpn := ConvRPN(input)
		fmt.Println("RPN: ", rpn)
		result := Calc(rpn)
		fmt.Println("Result: ", result)
		fmt.Println()
	}
}

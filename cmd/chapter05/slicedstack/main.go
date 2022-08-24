package main

import (
	"fmt"
	"github.com/ahamtat/go-generic-algorithms/pkg/stack"
)

func main() {
	// Create a stack of names
	nameStack := stack.Sliced[string]{}
	nameStack.Push("Zachary")
	nameStack.Push("Adolf")

	if !nameStack.IsEmpty() {
		fmt.Printf("\nTop of stack is %s", nameStack.Top())
	}

	if !nameStack.IsEmpty() {
		fmt.Printf("\nValue popped from stack is %s", nameStack.Pop())
	}

	if !nameStack.IsEmpty() {
		fmt.Printf("\nValue popped from stack is %s", nameStack.Pop())
	}

	if !nameStack.IsEmpty() {
		fmt.Printf("\nValue popped from stack is %s", nameStack.Pop())
	}

	if !nameStack.IsEmpty() {
		fmt.Printf("\nValue popped from stack is %s", nameStack.Pop())
	}

	// Create a stack of integers
	intStack := stack.Sliced[int]{}
	intStack.Push(5)
	intStack.Push(10)
	intStack.Push(0) // Problem since 0 is the zero

	// value for int
	top := intStack.Top()
	if !intStack.IsEmpty() {
		fmt.Printf("\nValue on top of intStack is %d", top)
	}

	if !intStack.IsEmpty() {
		fmt.Printf("\nValue popped from intStack is %d", intStack.Pop())
	}

	if !intStack.IsEmpty() {
		fmt.Printf("\nValue popped from intStack is %d", intStack.Pop())
	}

	if !intStack.IsEmpty() {
		fmt.Printf("\nValue popped from intStack is %d", intStack.Pop())
	}
}

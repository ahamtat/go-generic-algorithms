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

	topOfStack := nameStack.Top()
	if topOfStack != stack.GetZero[string]() {
		fmt.Printf("\nTop of stack is %s", topOfStack)
	}

	poppedFromStack := nameStack.Pop()
	if poppedFromStack != stack.GetZero[string]() {
		fmt.Printf("\nValue popped from stack is %s",
			poppedFromStack)
	}

	poppedFromStack = nameStack.Pop()
	if poppedFromStack != stack.GetZero[string]() {
		fmt.Printf("\nValue popped from stack is %s",
			poppedFromStack)
	}

	poppedFromStack = nameStack.Pop()
	if poppedFromStack != stack.GetZero[string]() {
		fmt.Printf("\nValue popped from stack is %s",
			poppedFromStack)
	}

	poppedFromStack = nameStack.Pop()
	if poppedFromStack != stack.GetZero[string]() {
		fmt.Printf("\nValue popped from stack is %s",
			poppedFromStack)
	}

	// Create a stack of integers
	intStack := stack.Sliced[int]{}
	intStack.Push(5)
	intStack.Push(10)
	intStack.Push(0) // Problem since 0 is the zero

	// value for int
	top := intStack.Top()
	if top != stack.GetZero[int]() {
		fmt.Printf("\nValue on top of intStack is %d", top)
	}

	popFromStack := intStack.Pop()
	if popFromStack != stack.GetZero[int]() {
		fmt.Printf("\nValue popped from intStack is %d", popFromStack)
	}

	popFromStack = intStack.Pop()
	if popFromStack != stack.GetZero[int]() {
		fmt.Printf("\nValue popped from intStack is %d", popFromStack)
	}

	popFromStack = intStack.Pop()
	if popFromStack != stack.GetZero[int]() {
		fmt.Printf("\nValue popped from intStack is %d", popFromStack)
	}
}

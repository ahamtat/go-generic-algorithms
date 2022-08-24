package stack

import "github.com/ahamtat/go-generic-algorithms/pkg/data"

type Sliced[T data.Ordered] struct {
	items []T
}

func GetZero[T data.Ordered]() T {
	var result T
	return result
}

// Methods
func (stack *Sliced[T]) Push(item T) {
	// item is added to the right-most position in the
	// slice
	if item != GetZero[T]() { // We exclude item if it
		// is GetZero[T]()
		stack.items = append(stack.items, item)
	}
}

func (stack *Sliced[T]) Pop() T {
	length := len(stack.items)
	if length > 0 {
		returnValue := stack.items[length-1]
		stack.items = stack.items[:(length - 1)]
		return returnValue
	} else {
		return GetZero[T]()
	}
}

func (stack Sliced[T]) Top() T {
	length := len(stack.items)
	if length > 0 {
		return stack.items[length-1]
	} else {
		return GetZero[T]()
	}
}

func (stack Sliced[T]) IsEmpty() bool {
	return len(stack.items) == 0
}

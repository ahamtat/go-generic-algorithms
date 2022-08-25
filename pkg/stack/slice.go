package stack

type Sliced[T any] struct {
	items []T
}

// Methods
func (stack *Sliced[T]) Push(item T) {
	// item is added to the right-most position in the slice
	stack.items = append(stack.items, item)
}

func (stack *Sliced[T]) Pop() (T, error) {
	var result T

	if stack.IsEmpty() {
		return result, ErrEmptyStack
	}

	length := len(stack.items)
	result = stack.items[length-1]
	stack.items = stack.items[:(length - 1)]

	return result, nil
}

func (stack Sliced[T]) Top() T {
	length := len(stack.items)
	return stack.items[length-1]
}

func (stack Sliced[T]) IsEmpty() bool {
	return len(stack.items) == 0
}

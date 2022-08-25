package stack

type Node[T any] struct {
	value T
	next  *Node[T]
}

type Nodes[T any] struct {
	first *Node[T]
}

func (n *Nodes[T]) Push(item T) {
	newNode := Node[T]{item, n.first}
	n.first = &newNode
}

func (n *Nodes[T]) Top() T {
	return n.first.value
}

func (n *Nodes[T]) Pop() (T, error) {
	var result T

	if n.IsEmpty() {
		return result, ErrEmptyStack
	}

	result = n.Top()
	n.first = n.first.next

	return result, nil
}

func (n *Nodes[T]) IsEmpty() bool {
	return n.first == nil
}

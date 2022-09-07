package queue

type Node[T any] struct {
	item T
	next *Node[T]
}

type Nodes[T any] struct {
	first, last *Node[T]
	length      int
}

type NodesIterator[T any] struct {
	next *Node[T]
}

func (queue *Nodes[T]) Insert(item T) {
	newNode := &Node[T]{item, nil}

	if queue.first == nil {
		queue.first = newNode
		queue.last = queue.first
	} else {
		queue.last.next = newNode
		queue.last = newNode
	}

	queue.length++
}

func (queue *Nodes[T]) Remove() T {
	returnValue := queue.first.item
	queue.first = queue.first.next

	if queue.first == nil {
		queue.last = nil
	}

	return returnValue
}

func (queue *Nodes[T]) First() T {
	return queue.first.item
}

func (queue *Nodes[T]) Size() int {
	return queue.length
}

func (queue *Nodes[T]) Range() NodesIterator[T] {
	return NodesIterator[T]{queue.first}
}

func (NodesIterator *NodesIterator[T]) Empty() bool {
	return NodesIterator.next == nil
}

func (NodesIterator *NodesIterator[T]) Next() T {
	returnValue := NodesIterator.next.item

	if NodesIterator.next != nil {
		NodesIterator.next = NodesIterator.next.next
	}

	return returnValue
}

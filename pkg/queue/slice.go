package queue

type Sliced[T any] struct {
	items []T
}

// Sliced Methods

func (queue *Sliced[T]) Insert(item T) {
	// item is added to the right-most position in the slice
	queue.items = append(queue.items, item)
}

func (queue *Sliced[T]) Remove() T {
	returnValue := queue.items[0]
	queue.items = queue.items[1:]
	return returnValue
}

func (queue *Sliced[T]) First() T {
	return queue.items[0]
}

func (queue *Sliced[T]) Size() int {
	return len(queue.items)
}

func (queue *Sliced[T]) Range() Iterator[T] {
	return Iterator[T]{0, queue.items}
}

type Iterator[T any] struct {
	next  int // index in items
	items []T
}

// Iterator Methods

func (iterator *Iterator[T]) Empty() bool {
	return iterator.next == len(iterator.items)
}

func (iterator *Iterator[T]) Next() T {
	returnValue := iterator.items[iterator.next]
	iterator.next++
	return returnValue
}

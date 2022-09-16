package dequeue

type Dequeue[T any] interface {
	InsertFront(item T)
	InsertBack(item T)
	First() T
	RemoveFirst() T
	Last() T
	RemoveLast() T
	Empty() bool
}

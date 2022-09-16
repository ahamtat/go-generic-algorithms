package dequeue

type Sliced[T any] struct {
	items []T
}

func (s *Sliced[T]) InsertFront(item T) {
	s.items = append(s.items, item) // Expands deque.items
	for i := len(s.items) - 1; i > 0; i-- {
		s.items[i] = s.items[i-1]
	}
	s.items[0] = item
}

func (s *Sliced[T]) InsertBack(item T) {
	s.items = append(s.items, item)
}

func (s *Sliced[T]) First() T {
	return s.items[0]
}

func (s *Sliced[T]) RemoveFirst() T {
	returnValue := s.items[0]
	s.items = s.items[1:]
	return returnValue
}

func (s *Sliced[T]) Last() T {
	return s.items[len(s.items)-1]
}

func (s *Sliced[T]) RemoveLast() T {
	length := len(s.items)
	returnValue := s.items[length-1]
	s.items = s.items[:(length - 1)]
	return returnValue
}

func (s *Sliced[T]) Empty() bool {
	return len(s.items) == 0
}

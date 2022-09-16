package dequeue

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSliced_InsertFront(t *testing.T) {
	d := &Sliced[int]{}
	insertFrontDequeueItem[int](d, 10)
	require.False(t, d.Empty())
}

func insertFrontDequeueItem[T any](d Dequeue[T], item T) {
	d.InsertFront(item)
}

package queue

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type nodesTestCases[T any] struct {
	description string
	object      Sliced[T]
	input       []T
	expected    T
}

func TestNodes_Insert(t *testing.T) {
	for _, scenario := range []nodesTestCases[int]{
		{
			"insert on empty queue",
			Sliced[int]{},
			[]int{},
			0,
		},
		{
			"insert on valid queue",
			Sliced[int]{},
			[]int{15, 20, 30},
			3,
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			for _, item := range scenario.input {
				scenario.object.Insert(item)
			}
			require.Equal(t, scenario.expected, scenario.object.Size())
		})
	}
}

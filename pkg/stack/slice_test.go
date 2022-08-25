package stack

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type slicedTestCases[T any] struct {
	description string
	object      Sliced[T]
	input       []T
	expected    T
	expectedErr error
}

func TestSliced_Pop(t *testing.T) {
	for _, scenario := range []slicedTestCases[int]{
		{
			"pop from empty int stack",
			Sliced[int]{},
			[]int{},
			0,
			ErrEmptyStack,
		},
		{
			"populate int stack",
			Sliced[int]{},
			[]int{18, 22},
			22,
			nil,
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			// populate object with items
			for _, item := range scenario.input {
				scenario.object.Push(item)
			}
			value, err := scenario.object.Pop()
			require.Equal(t, scenario.expected, value)
			require.Equal(t, scenario.expectedErr, err)
		})
	}
}

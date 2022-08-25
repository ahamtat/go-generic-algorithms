package stack

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type nodeTestCases[T any] struct {
	description string
	object      Nodes[T]
	input       []T
	expected    T
	expectedErr error
}

func TestNodes_Pop(t *testing.T) {
	for _, scenario := range []nodeTestCases[string]{
		{
			"pop from empty string stack",
			Nodes[string]{},
			[]string{},
			"",
			ErrEmptyStack,
		},
		{
			"populate string stack",
			Nodes[string]{},
			[]string{"Beatrice", "Veronica"},
			"Veronica",
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

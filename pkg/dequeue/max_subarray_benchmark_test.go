package dequeue

import (
	"math/rand"
	"testing"
)

const (
	arraySize    = 1_000_000
	subarraySize = 3
)

func BenchmarkBruteForce_MaxSubArray_Int(b *testing.B) {
	input := createInputDataInt(arraySize)

	for i := 0; i < b.N; i++ {
		//
		for first := 0; first <= len(input)-subarraySize; first++ {
			max := input[first]
			for second := 0; second < subarraySize; second++ {
				if input[first+second] > max {
					max = input[first+second]
				}
			}
		}
	}
}

func BenchmarkSliced_MaxSubArray_Int(b *testing.B) {
	input := createInputDataInt(arraySize)

	for i := 0; i < b.N; i++ {
		maxSubarrayUsingDeque(input, subarraySize)
	}
}

func createInputDataInt(size int) []int {
	result := make([]int, size)
	for i := 0; i < size; i++ {
		result[i] = rand.Intn(10000)
	}
	return result
}

func maxSubarrayUsingDeque(input []int, k int) {
	deque := Sliced[int]{}
	var index int

	// First window
	for index = 0; index < k; index++ {
		for {
			if deque.Empty() || input[index] < input[deque.Last()] {
				break
			}
			deque.RemoveLast()
		}
		deque.InsertBack(index)
	}

	for ; index < len(input); index++ {
		// Remove elements out of the window
		for {
			if deque.Empty() || deque.First() > index-k {
				break
			}
			deque.RemoveFirst()
		}

		// Remove values smaller than the element currently being added
		for {
			if deque.Empty() || input[index] < input[deque.Last()] {
				break
			}
			deque.RemoveLast()
		}
		deque.InsertBack(index)
	}
}

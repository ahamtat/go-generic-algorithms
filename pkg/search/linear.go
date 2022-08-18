package search

import (
	"github.com/ahamtat/go-generic-algorithms/pkg/data"
	"runtime"
)

func Linear[T data.Ordered](slice []T, target T) bool {
	// Return true if T is in the slice
	for i := 0; i < len(slice); i++ {
		if slice[i] == target {
			return true
		}
	}
	return false
}

func searchSegment[T data.Ordered](data []T, target T, low, high int, ch chan<- bool) {
	for i := low; i < high; i++ {
		if data[i] == target {
			ch <- true
		}
	}
	ch <- false
}

func ConcurrentLinear[T data.Ordered](data []T, target T) bool {
	ch := make(chan bool)
	numSegments := runtime.NumCPU()
	segmentSize := int(float64(len(data)) / float64(numSegments))

	for i := 0; i < numSegments; i++ {
		go searchSegment(data, target, i*segmentSize, (i+1)*segmentSize, ch)
	}

	var completed int
OUTER:
	for {
		if <-ch {
			return true
		}
		completed++

		if completed == numSegments {
			break OUTER
		}
	}
	return false
}

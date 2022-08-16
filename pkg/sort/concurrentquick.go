package sort

import (
	"github.com/ahamtat/go-generic-algorithms/pkg/data"
	"sync"
)

const (
	minLength = 30
	threshold = 5000
)

func middlePartition[T data.Ordered](data []T) int {
	data[len(data)/2], data[0] = data[0], data[len(data)/2]
	pivot := data[0]
	mid := 0
	i := 1
	for i < len(data) {
		if data[i] < pivot {
			mid += 1
			data[i], data[mid] = data[mid], data[i]
		}
		i += 1
	}
	data[0], data[mid] = data[mid], data[0]
	return mid
}

func ConcurrentQuicksort[T data.Ordered](data []T, wg *sync.WaitGroup) {
	for len(data) >= minLength {
		mid := middlePartition(data)

		var portion []T
		if mid < len(data)/2 {
			portion = data[:mid]
			data = data[mid+1:]
		} else {
			portion = data[mid+1:]
			data = data[:mid]
		}

		if len(portion) > threshold {
			wg.Add(1)
			go func(data []T) {
				defer wg.Done()
				ConcurrentQuicksort(data, wg)
			}(portion)
		} else {
			ConcurrentQuicksort(portion, wg)
		}
	}

	InsertSort(data)
}

func QSort[T data.Ordered](data []T) {
	var wg sync.WaitGroup
	ConcurrentQuicksort[T](data, &wg)
	wg.Wait()
}

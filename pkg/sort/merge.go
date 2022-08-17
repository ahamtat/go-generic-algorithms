package sort

import (
	"github.com/ahamtat/go-generic-algorithms/pkg/data"
	"sync"
)

const maxLength = 5000

func merge[T data.Ordered](left, right []T) []T {
	result := make([]T, len(left)+len(right))

	var i, j, k int
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		result[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		result[k] = right[j]
		j++
		k++
	}

	return result
}
func Merge[T data.Ordered](data []T) []T {
	if len(data) > 100 {
		middle := len(data) / 2
		left := data[:middle]
		right := data[middle:]
		data = merge(Merge(right), Merge(left))
	} else {
		InsertSort(data)
	}
	return data
}

func ConcurrentMerge[T data.Ordered](data []T) []T {
	if len(data) > 1 {
		if len(data) <= maxLength {
			return Merge(data)
		} else { // Concurrent
			middle := len(data) / 2
			left := data[:middle]
			right := data[middle:]

			var wg sync.WaitGroup
			wg.Add(2)

			var data1, data2 []T
			go func() {
				defer wg.Done()
				data1 = ConcurrentMerge(left)
			}()
			go func() {
				defer wg.Done()
				data2 = ConcurrentMerge(right)
			}()

			wg.Wait()
			return merge(data1, data2)
		}
	}
	return nil
}

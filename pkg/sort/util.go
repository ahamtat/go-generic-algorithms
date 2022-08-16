package sort

import "github.com/ahamtat/go-generic-algorithms/pkg/data"

func IsSorted[T data.Ordered](data []T) bool {
	for i := 1; i < len(data); i++ {
		if data[i] < data[i-1] {
			return false
		}
	}
	return true
}

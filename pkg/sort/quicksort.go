package sort

import "github.com/ahamtat/go-generic-algorithms/pkg/data"

func Quick[T data.Ordered](data []T, low, high int) {
	if low < high {
		pivot := partition(data, low, high)
		Quick(data, low, pivot)
		Quick(data, pivot+1, high)
	}
}

func partition[T data.Ordered](data []T, low, high int) int {
	// Pick a lowest bound element as a pivot value
	pivot := data[low]

	i, j := low, high
	for i < j {
		for data[i] <= pivot && i < high {
			i++
		}
		for data[j] > pivot && j > low {
			j--
		}
		if i < j {
			data[i], data[j] = data[j], data[i]
		}
	}
	data[low], data[j] = data[j], pivot

	return j
}

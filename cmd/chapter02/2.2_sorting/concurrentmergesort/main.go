package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ahamtat/go-generic-algorithms/pkg/sort"
)

const size = 50_000_000

func main() {
	data := make([]float64, size)
	for i := 0; i < size; i++ {
		data[i] = 100.0 * rand.Float64()
	}

	start := time.Now()
	result := sort.ConcurrentMerge(data)
	elapsed := time.Since(start)
	fmt.Println("Elapsed time for concurrent mergesort = ", elapsed)
	fmt.Println("Sorted: ", sort.IsSorted(result))

	start = time.Now()
	result = sort.Merge[float64](data)
	elapsed = time.Since(start)
	fmt.Println("Elapsed time for MergeSort = ", elapsed)
	fmt.Println("Is sorted: ", sort.IsSorted(result))
}

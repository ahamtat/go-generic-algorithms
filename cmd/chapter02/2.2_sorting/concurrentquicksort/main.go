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

	data2 := make([]float64, size)
	copy(data2, data)

	start := time.Now()
	sort.QSort[float64](data)
	elapsed := time.Since(start)
	fmt.Println("Elapsed time for concurrent quicksort = ", elapsed)
	fmt.Println("Is sorted: ", sort.IsSorted(data))

	start = time.Now()
	sort.Quick(data2, 0, len(data2)-1)
	elapsed = time.Since(start)
	fmt.Println("Elapsed time for regular quicksort = ", elapsed)
	fmt.Println("Is sorted: ", sort.IsSorted(data2))
}

package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"

	"github.com/ahamtat/go-generic-algorithms/pkg/search"
)

const size = 100_000_000

func main() {
	fmt.Printf("CPU number: %d\n", runtime.NumCPU())

	data := make([]float64, size)
	for i := 0; i < size; i++ {
		data[i] = 100.0 * rand.Float64()
	}

	start := time.Now()
	result := search.Linear[float64](data, 54.0)
	elapsed := time.Since(start)
	fmt.Println("Time to search slice of 100_000_000 floats using linearSearch = ", elapsed)
	fmt.Println("Result of search is ", result)

	start = time.Now()
	result = search.Linear[float64](data, data[size/2])
	elapsed = time.Since(start)
	fmt.Println("Time to search slice of 100_000_000 floats using linearSearch = ", elapsed)
	fmt.Println("Result of search is ", result)

	start = time.Now()
	result = search.ConcurrentLinear[float64](data, 54.0) // Should return false
	elapsed = time.Since(start)
	fmt.Println("Time to search slice of 100_000_000 floats using concurrentSearch = ", elapsed)
	fmt.Println("Result of search is ", result)

	start = time.Now()
	result = search.ConcurrentLinear[float64](data, data[size/2]) // true
	elapsed = time.Since(start)
	fmt.Println("Time to search slice of 100_000_000 floats using concurrentSearch = ", elapsed)
	fmt.Println("Result of search is ", result)
}

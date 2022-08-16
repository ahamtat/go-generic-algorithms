package main

import (
	"fmt"
	"math"
	"time"

	"github.com/ahamtat/go-generic-algorithms/pkg/sort"
)

const size = 100_000

func main() {
	data := make([]float64, size)
	for i := 0; i < size; i++ {
		data[i] = math.Sin(float64(i * i))
	}

	start := time.Now()
	sort.Quick[float64](data, 0, len(data)-1)
	elapsed := time.Since(start)
	fmt.Println("Elapsed sort time for sine wave using quicksort: ", elapsed)

	data = make([]float64, size)
	for i := 0; i < size; i++ {
		data[i] = math.Sin(float64(i * i))
	}

	start = time.Now()
	sort.Bubble[float64](data)
	elapsed = time.Since(start)
	fmt.Println("Elapsed sort time for sine wave using bubblesort: ", elapsed)
}

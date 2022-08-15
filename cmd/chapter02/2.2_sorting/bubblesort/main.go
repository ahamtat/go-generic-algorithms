package main

import (
	"fmt"

	"github.com/ahamtat/go-generic-algorithms/pkg/sort"
)

func main() {
	numbers := []float64{3.5, -2.4, 12.8, 9.1}
	names := []string{"Zachary", "John", "Moe", "Jim", "Robert"}

	sort.Bubble[float64](numbers)
	fmt.Println(numbers)

	sort.Bubble[string](names)
	fmt.Println(names)
}

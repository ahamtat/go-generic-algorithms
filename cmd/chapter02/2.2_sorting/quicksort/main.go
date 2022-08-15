package main

import (
	"fmt"

	"github.com/ahamtat/go-generic-algorithms/pkg/sort"
)

func main() {
	numbers := []float64{3.5, -2.4, 12.8, 9.1}
	names := []string{"Zachary", "John", "Moe", "Jim", "Robert"}

	sort.Quick[float64](numbers, 0, len(numbers)-1)
	fmt.Println(numbers)

	sort.Quick[string](names, 0, len(names)-1)
	fmt.Println(names)
}

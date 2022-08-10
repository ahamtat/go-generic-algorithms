package main

import (
	"fmt"
	"time"
)

const LargestPrime = 10_000_000

func SieveOfEratosthenes(n int) []int {
	// Finds all primes up to n
	primes := make([]bool, n+1)
	for i := 2; i < n+1; i++ {
		primes[i] = true
	}

	// The Sieve logic
	for p := 2; p*p <= n; p++ {
		if primes[p] == true {
			// Update all multiples of p
			for i := p * 2; i <= n; i += p {
				primes[i] = false
			}
		}
	}

	// return all prime numbers <= n
	var primeNumbers []int
	for p := 2; p <= n; p++ {
		if primes[p] == true {
			primeNumbers = append(primeNumbers, p)
		}
	}
	return primeNumbers
}

func main() {
	start := time.Now()
	sieve := SieveOfEratosthenes(LargestPrime)
	elapsed := time.Since(start)
	fmt.Println("\nComputation time: ", elapsed)
	fmt.Println("Largest prime: ", sieve[len(sieve)-1])
}

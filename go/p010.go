package main

import (
	"fmt"
	"math"
)

/*
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
*/

func sieveOfEratosthenes(n int64) []bool {
	a := make([]bool, n)
	for i := int64(2); i < n; i++ {
		a[i] = true
	}
	sqrt := int64(math.Sqrt(float64(n)))
	for i := int64(2); i < sqrt; i++ {
		if a[i] {
			tmp := i * i
			for j, c := tmp, int64(0); j < n; c, j = c+1, tmp+(c*i) {
				a[j] = false
			}
		}
	}
	return a
}

func main() {
	const max = 2e6 //Maximum two million
	primes := sieveOfEratosthenes(max)
	var sum int64
	for i := range primes {
		if primes[i] {
			sum += int64(i)
		}
	}
	fmt.Printf("Sum = %d\n", sum) //outputs: 142913828922
}

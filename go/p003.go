package main

import "fmt"

/*
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/

func main() {
	const number = 600851475143
	var factors []int64
	d := int64(2)
	i := int64(number)
	for i > 1 {
		for i%d == 0 {
			factors = append(factors, d)
			i /= d
		}
		d++
	}
	fmt.Printf("Largest Factor = %d\n", factors[len(factors)-1]) //outputs: 6857
}

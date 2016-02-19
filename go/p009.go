package main

import "fmt"

/*
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
a2 + b2 = c2

For example, 32 + 42 = 9 + 16 = 25 = 52.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
*/

func main() {
	const max = int32(1e3)
	var product int32
	for a := int32(1); a < int32(max/3); a++ {
		for b := a + 1; b < (max - b - a); b++ {
			c := max - b - a
			if (a*a)+(b*b) == (c * c) {
				product = a * b * c
				break
			}
		}
	}
	fmt.Printf("Product = %d\n", product) //output: 31875000
}

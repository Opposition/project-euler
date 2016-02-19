package main

import "fmt"

/*
If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
*/

func isMultiple(n int32) bool {
	return n%3 == 0 || n%5 == 0
}

func main() {
	var sum int32
	for i := int32(1); i < 1000; i++ {
		if isMultiple(i) {
			sum += i
		}
	}
	fmt.Printf("Sum = %d\n", sum) //outputs: 233168
}

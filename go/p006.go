package main

import (
	"fmt"
	"math"
)

/*
The sum of the squares of the first ten natural numbers is: 1^2 + 2^2 + ... + 10^2 = 385

The square of the sum of the first ten natural numbers is: (1 + 2 + ... + 10)^2 = 55^2 = 3025

Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
*/

func main() {
	var sumOfSquare float64
	var squareOfSum float64
	for i := int64(1); i <= 1e2; i++ {
		sumOfSquare += math.Pow(float64(i), 2)
		squareOfSum += float64(i)
	}
	squareOfSum = math.Pow(float64(squareOfSum), 2)
	result := int64(math.Abs(squareOfSum - sumOfSquare))
	fmt.Printf("Difference = %d\n", result) //outputs: 25164150
}

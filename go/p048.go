package main

import (
	"fmt"
	"math/big"
)

/*

The series, 1^1 + 2^2 + 3^3 + ... + 10^10 = 10405071317.

Find the last ten digits of the series, 1^1 + 2^2 + 3^3 + ... + 1000^1000.

*/

func pow(n, p int) *big.Int {
	constNum := big.NewInt(int64(n))
	result := big.NewInt(1)
	for i := 0; i < p; i++ {
		result.Mul(result, constNum)
	}
	return result
}

func main() {
	result := new(big.Int)
	for i := 1; i <= 1000; i++ {
		result.Add(result, pow(i, i))
	}
	fmt.Printf("Result = %s\n", result.String()[len(result.String())-10:]) //outputs: 9110846700
}

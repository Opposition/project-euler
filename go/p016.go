package main

import (
	"fmt"
	"math/big"
)

/*

2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 2^1000?

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
	p := pow(2, 1000)
	var sum int64
	for i := 0; i < len(p.String()); i++ {
		sum += int64(rune(p.String()[i]) - '0')
	}
	fmt.Printf("Sum = %d\n", sum) //ouputs: 1366
}

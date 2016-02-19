package main

import (
	"fmt"
	"math/big"
)

/*
n! means n × (n − 1) × ... × 3 × 2 × 1

For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

Find the sum of the digits in the number 100!
*/

func main() {
	fac := big.NewInt(1)
	for i := int64(100); i >= 1; i-- {
		fac = fac.Mul(fac, big.NewInt(i))
	}
	var sum int32
	k := fac.String()
	for i := 0; i < len(k); i++ {
		sum += int32(rune(k[i]) - '0')
	}
	fmt.Printf("Sum = %d\n", sum) //outputs: 648
}

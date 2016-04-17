package main

import (
	"fmt"
	"strconv"
)

/*

The number 3797 has an interesting property. Being prime itself, it is possible to continuously remove digits from left to right, and remain prime at each stage: 3797, 797, 97, and 7. Similarly we can work from right to left: 3797, 379, 37, and 3.

Find the sum of the only eleven primes that are both truncatable from left to right and right to left.

NOTE: 2, 3, 5, and 7 are not considered to be truncatable primes.

*/

func isPrime(n int64) bool {
	if n&1 == 0 {
		return n == 2
	}
	for i := int64(3); (i * i) <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return n != 1
}

func leftToRight(n int64) bool {
	m := strconv.FormatInt(n, 10)
	for i := 0; i < len(m); i++ {
		digit, _ := strconv.ParseInt(m[i:], 10, 64)
		if !isPrime(digit) {
			return false
		}
	}
	return true
}

func rightToLeft(n int64) bool {
	m := strconv.FormatInt(n, 10)
	for i := len(m) - 1; i >= 0; i-- {
		digit, _ := strconv.ParseInt(m[:i+1], 10, 64)
		if !isPrime(digit) {
			return false
		}
	}
	return true
}

func main() {
	var sum int64
	//NOTE: 2, 3, 5, and 7 are not considered to be truncatable primes. So we start at the next up Prime, which is 11
	for i, j := int64(11), int64(0); j != 11; i++ {
		if leftToRight(i) && rightToLeft(i) {
			sum += i
			j++
		}
	}
	fmt.Printf("Sum = %d\n", sum) //outputs: 748317
}

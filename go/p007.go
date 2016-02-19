package main

import "fmt"

/*
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?
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

func main() {
	var value, index int32
	for {
		if isPrime(int64(value)) {
			index++
			if index == 10001 {
				break
			}
		}
		value++
	}
	fmt.Printf("The answer = %d\n", value) //output: 104743
}

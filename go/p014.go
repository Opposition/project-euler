package main

import "fmt"

/*

The following iterative sequence is defined for the set of positive integers:

n → n/2 (n is even)
n → 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:

13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.

*/

func calculate(n int64) int64 {
	if n&1 == 0 {
		return n / 2
	}
	return 3*n + 1
}

func gauntlet(n int64) int64 {
	chain := int64(1)
	for n != 1 {
		chain++
		n = calculate(n)
	}
	return chain
}

func main() {
	var highestChain int64
	var highestNumber int64
	for i := int64(1); i < 1e6; i++ {
		chain := gauntlet(i)
		if chain > highestChain {
			highestChain = chain
			highestNumber = i
		}
	}
	fmt.Printf("Number under one million that produces the longest chain = %d\n", highestNumber) //output: 837799
}

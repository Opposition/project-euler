package main

import "fmt"

/*
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/

func isEvenlyDivisible(n int32) bool {
	for i := int32(1); i <= 20; i++ {
		if n%i != 0 {
			return false
		}
	}
	return true
}

func main() {
	i := int32(1)
	for {
		//making sure it is a even number, before we call the function, shaves 1 second off the overall time.
		if i&1 == 0 && isEvenlyDivisible(i) {
			break
		}
		i++
	}
	fmt.Printf("Smallest positive number = %d\n", i) //outputs: 232792560
}

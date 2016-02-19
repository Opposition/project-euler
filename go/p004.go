package main

import (
	"fmt"
	"strconv"
)

/*
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/

func isPalindrome(n int32) bool {
	text := strconv.Itoa(int(n))
	for i, j := 0, len(text)-1; i < len(text)-1; i, j = i+1, j-1 {
		if text[i] != text[j] {
			return false
		}
	}
	return true
}

func main() {
	var largestPalindrome int32
	for x := int32(1e2); x < 1e3; x++ {
		for y := int32(1e2); y < 1e3; y++ {
			number := x * y
			if isPalindrome(number) {
				if number > largestPalindrome {
					largestPalindrome = number
				}
			}
		}
	}
	fmt.Printf("Largest Palindrome = %d\n", largestPalindrome) //outputs: 906609
}

package main

import "fmt"

func isPalindromeByteArray(n []byte) bool {
	for i, j := 0, len(n)-1; i < len(n)-1; i, j = i+1, j-1 {
		if n[i] != n[j] {
			return false
		}
	}
	return true
}

func isPalindromeNumber(n int32) bool {
	x := n
	num := int32(0)
	for x > 0 {
		num = num*10 + (x % 10)
		x /= 10
	}
	return num == n
}

func getBinary(n int32, buf []byte) (written int) {
	//we receive buffer to avoid heap allocation
	for i := len(buf) - 1; n != 0; i-- {
		buf[i] = byte(n & 1)
		n >>= 1
		written++
	}
	return
}

func main() {
	buf := make([]byte, 31) //32 bit since we are using int32
	var sum int32
	for i := int32(0); i < 1e6; i++ {
		w := getBinary(i, buf)
		if isPalindromeNumber(i) && isPalindromeByteArray(buf[len(buf)-w:]) {
			sum += i
		}
	}
	fmt.Printf("Sum = %d\n", sum) //outputs: 872187
	//finishes in about 40-60 ms. Deviated greatly from my first attempt, since the first one did too many allocations (700 ms).
}

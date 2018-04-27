package main

import (
	"fmt"
)

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func FindConsecutiveN(a, b int) int {
	isPrime := true
	n := 0
	for isPrime {
		result := n*n + a*n + b
		isPrime = IsPrime(result)
		if isPrime {
			n++
		} else {
			return n
		}
	}
	return -1
}

func main() {
	maxN := 0
	maxA := -11111
	maxB := -11111
	for a := -999; a < 1000; a++ {
		for b := -999; b < 1000; b++ {
			n := FindConsecutiveN(a, b)
			if n == -1 {
				panic("n cannot be -1")
			}
			if n > maxN {
				maxN = n
				maxA = a
				maxB = b
				fmt.Println("current max (a,b,n)", a, b, n)
			}
		}
	}
	fmt.Println("max (a,b,n)", maxA, maxB, maxN)
}

package main

import (
	"fmt"
	"strconv"
)

func IsFifthPowerEqual(n int) bool {
	s := strconv.Itoa(n)
	sum := 0
	for _, digit := range s {
		d, err := strconv.Atoi(string(digit))
		if err != nil {
			panic(err)
		}
		sum += d * d * d * d * d
	}
	if sum == n {
		return true
	}
	return false
}

func main() {
	d95 := 9 * 9 * 9 * 9 * 9
	fmt.Println("9^5 =", d95)
	fmt.Println("max possible: 6 * 9^5 =", 6*d95)
	sum := 0
	for i := 2; i < 6*d95; i++ {
		if IsFifthPowerEqual(i) {
			fmt.Println(i)
			sum += i
		}
	}
	fmt.Println("Sum:", sum)
}

package lychrel

import (
	"fmt"
	"strconv"
)

func Reverse(s string) (rs string) {
	for _, r := range s {
		rs = string(r) + rs
	}
	return
}

func IsDecimalPalindromeNumber(n int64) bool {
	if n < 0 {
		n = -n
	}

	s := strconv.FormatInt(n, 10)
	bound := (len(s) / 2) + 1
	for i := 0; i < bound; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func LychrelNumberTest(n int64, maxIteration int) bool {
	s := strconv.FormatInt(n, 10)
	if len(s) < 2 {
		panic("input number must be at least 2 digits!")
	}
	if maxIteration < 0 {
		panic("negative iteration number")
	}

	fmt.Println("Lychrel Number Test", n)
	nextN := n
	for i := 0; i < maxIteration; i++ {
		// reverse nextN
		nstr := strconv.FormatInt(nextN, 10)
		rnstr := Reverse(nstr)
		rn, _ := strconv.ParseInt(rnstr, 10, 64)
		// add the reverse back to nextN
		nextN = nextN + rn

		fmt.Println("iteration", i+1)
		fmt.Println(nextN)
		if IsDecimalPalindromeNumber(nextN) {
			return true
		}
	}
	return false
}

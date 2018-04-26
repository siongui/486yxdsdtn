package db3

import (
	"strconv"
)

func IsIntDivisibleBy3(n int) bool {
	digits := strconv.Itoa(n)
	sumOfDigits := 0
	for _, digit := range digits {
		d, _ := strconv.Atoi(string(digit))
		sumOfDigits += d
	}

	return (sumOfDigits % 3) == 0
}

func IsStrDivisibleBy3(n string) (bool, error) {
	sumOfDigits := 0
	for _, digit := range n {
		d, err := strconv.Atoi(string(digit))
		if err != nil {
			return false, err
		}
		sumOfDigits += d
	}

	return (sumOfDigits % 3) == 0, nil
}

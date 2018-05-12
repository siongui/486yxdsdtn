package politeness

// Get all prime factors of a given number n
func PrimeFactors(n int) (pfs []int) {
	// Get the number of 2s that divide n
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		pfs = append(pfs, n)
	}

	return
}

// Algorithm from wiki: An easy way of calculating the politeness of a positive
// number is that of decomposing the number into its prime factors, taking the
// powers of all prime factors greater than 2, adding 1 to all of them,
// multiplying the numbers thus obtained with each other and subtracting 1.
func CalculatePoliteness(n int) int {
	pfs := PrimeFactors(n)

	// key: prime
	// value: prime exponent
	m := make(map[int]int)
	for _, prime := range pfs {
		_, ok := m[prime]
		if ok {
			m[prime] += 1
		} else {
			m[prime] = 1
		}
	}

	politeness := 1
	for prime, exponent := range m {
		if prime > 2 {
			politeness *= (exponent + 1)
		}
	}
	politeness -= 1

	return politeness
}

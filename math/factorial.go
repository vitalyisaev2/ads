package math

import "fmt"

// NaiveFactorial algorithm, O(N)
func NaiveFactorial(n uint64) uint64 {
	var val uint64 = 1
	for i := uint64(2); i <= n; i++ {
		val *= i
	}
	return val
}

// NaiveFactorialParallel utilises all available cores, O(N)
func NaiveFactorialParallel(n uint64, cores int) uint64 {
	if 2*n < uint64(cores) {
		return NaiveFactorial(n)
	}

	worker := func(start, end uint64, results chan uint64) {
		var val uint64
		if start == 0 {
			val = 1
		} else {
			val = start
		}
		for i := start + 1; i <= end; i++ {
			val *= i
		}
		results <- val
		fmt.Println(start, end, val)
	}

	// load every available core
	results := make(chan uint64)
	part := n / uint64(cores)
	for c := 0; c < cores-1; c++ {
		begin := uint64(c)*part + 1
		end := uint64(c+1) * part
		go worker(begin, end, results)
	}
	go worker(uint64(cores-1)*part+1, n, results)

	var overall uint64 = 1
	for c := 0; c < cores; c++ {
		result := <-results
		overall *= result
	}
	return overall
}

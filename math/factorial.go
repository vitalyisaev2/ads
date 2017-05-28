package math

import (
	"math/big"
)

// NaiveFactorial algorithm, O(N)
func NaiveFactorial(n *big.Int) *big.Int {
	one := big.NewInt(1)
	val := big.NewInt(1)

	for i := big.NewInt(2); i.Cmp(n) <= 0; i.Add(i, one) {
		val.Mul(val, i)
	}
	return val
}

// NaiveFactorialParallel utilises all available cores, O(N)
func NaiveFactorialParallel(n *big.Int, _cores int) *big.Int {
	var (
		tmp     big.Int
		one     = big.NewInt(1)
		results = make(chan *big.Int)
		cores   = big.NewInt(int64(_cores))
	)

	if tmp.Mul(big.NewInt(2), n).Cmp(cores) < 0 {
		return NaiveFactorial(n)
	}

	worker := func(start, end *big.Int, results chan *big.Int) {
		var val big.Int
		var i big.Int

		if start.Cmp(big.NewInt(0)) == 0 {
			val = *big.NewInt(1)
		} else {
			val = *start
		}

		i.Add(start, one)
		for ; i.Cmp(end) <= 0; i.Add(&i, one) {
			val.Mul(&val, &i)
		}
		results <- &val
	}

	// load every available core
	part := tmp.Div(n, cores)

	for c := 0; c < _cores-1; c++ {
		var start, end big.Int
		start.Mul(part, big.NewInt(int64(c))).Add(&start, one)
		end.Mul(part, big.NewInt(int64(c+1)))
		go worker(&start, &end, results)
	}

	var start big.Int
	start.Mul(big.NewInt(int64(_cores-1)), part).Add(&start, one)
	go worker(&start, n, results)

	overall := big.NewInt(1)
	for c := 0; c < _cores; c++ {
		overall.Mul(overall, <-results)
	}
	return overall
}

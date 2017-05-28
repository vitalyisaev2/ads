package math

import (
	"math/big"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NaiveFactorial(t *testing.T) {
	assert.Zero(t, big.NewInt(1).Cmp(NaiveFactorial(big.NewInt(0))))
	assert.Zero(t, big.NewInt(1).Cmp(NaiveFactorial(big.NewInt(1))))
	assert.Zero(t, big.NewInt(2).Cmp(NaiveFactorial(big.NewInt(2))))
	assert.Zero(t, big.NewInt(6).Cmp(NaiveFactorial(big.NewInt(3))))
	assert.Zero(t, big.NewInt(3628800).Cmp(NaiveFactorial(big.NewInt(10))))
}

func Test_NaiveFactorialParallel(t *testing.T) {
	cpu := runtime.NumCPU()
	assert.Zero(t, big.NewInt(1).Cmp(NaiveFactorialParallel(big.NewInt(0), cpu)))
	assert.Zero(t, big.NewInt(3628800).Cmp(NaiveFactorialParallel(big.NewInt(10), cpu)))
	assert.Zero(t, big.NewInt(6402373705728000).Cmp(NaiveFactorialParallel(big.NewInt(18), cpu)))
}

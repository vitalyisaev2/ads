package math

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NaiveFactorial(t *testing.T) {
	assert.Equal(t, uint64(1), NaiveFactorial(0))
	assert.Equal(t, uint64(1), NaiveFactorial(1))
	assert.Equal(t, uint64(2), NaiveFactorial(2))
	assert.Equal(t, uint64(6), NaiveFactorial(3))
	assert.Equal(t, uint64(3628800), NaiveFactorial(10))
}

func Test_NaiveFactorialParallel(t *testing.T) {
	assert.Equal(t, uint64(1), NaiveFactorialParallel(0, runtime.NumCPU()))
	assert.Equal(t, uint64(3628800), NaiveFactorialParallel(10, runtime.NumCPU()))
	assert.Equal(t, uint64(6402373705728000), NaiveFactorialParallel(18, runtime.NumCPU()))
}

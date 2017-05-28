package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RaiseNumberToNaturalPower(t *testing.T) {
	assert.Equal(t, 4, RaiseNumberToNaturalPower(2, 2))
	assert.Equal(t, 81, RaiseNumberToNaturalPower(3, 4))
	assert.Equal(t, 5764801, RaiseNumberToNaturalPower(7, 8))
}

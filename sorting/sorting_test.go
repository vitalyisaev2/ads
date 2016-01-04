package sorting

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Sortable array of integers
type NumericalArray []int

func (na NumericalArray) Len() int { return len(na) }

func (na NumericalArray) Less(i, j int) bool {
	return na[i] < na[j]
}

func (na NumericalArray) Swap(i, j int) {
	na[i], na[j] = na[j], na[i]
}

// Sortable string
type CharacterArray []rune

func (ca CharacterArray) Len() int { return len(ca) }

func (ca CharacterArray) Less(i, j int) bool {
	return ca[i] < ca[j]
}

func (ca CharacterArray) Swap(i, j int) {
	ca[i], ca[j] = ca[j], ca[i]
}

func NewCharacterArray(str string) CharacterArray {
	ca := []rune(str)
	return ca
}

func (ca CharacterArray) toString() string {
	return string(ca)
}

func TestSelectionSort(t *testing.T) {

	numericalArrayTests := []struct {
		input, expected NumericalArray
	}{
		{
			NumericalArray{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			NumericalArray{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			NumericalArray{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			NumericalArray{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			NumericalArray{-9, -8, -7, -6, -5, -4, -3, -2, -1, 0},
			NumericalArray{-9, -8, -7, -6, -5, -4, -3, -2, -1, 0},
		},
	}

	characterArrayTests := []struct {
		input, expected string
	}{
		{
			"0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!#$%&'()*+,-./:;?@[\\]^_`{|}~ \t\n\r\x0b\x0c",
			"\t\n\v\f\r !#$%&'()*+,-./0123456789:;?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~",
		},
		{
			"абвгдеёжзийклмнопрстуфхцчшщъыьэюя",
			"абвгдежзийклмнопрстуфхцчшщъыьэюяё",
		},
	}

	for _, test := range numericalArrayTests {
		SelectionSort(test.input)
		assert.Equal(t, test.input, test.expected)
	}
	for _, test := range characterArrayTests {
		input := NewCharacterArray(test.input)
		SelectionSort(input)
		assert.Equal(t, input.toString(), test.expected)
	}

}

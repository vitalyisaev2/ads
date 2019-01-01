package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMedianSortedArrays(t *testing.T) {
	type testcase struct {
		answer float64
		num1   []int
		num2   []int
	}

	cases := []testcase{
		{4.5, []int{1, 2, 3, 4}, []int{5, 6, 7, 8}},
		{4.5, []int{5, 6, 7, 8}, []int{1, 2, 3, 4}},
		{4.5, []int{2, 3, 4, 5}, []int{1, 6, 7, 8}},
		{4.5, []int{1, 6, 7, 8}, []int{2, 3, 4, 5}},
		{2.5, []int{1, 2}, []int{3, 4}},
		{2.0, []int{1, 3}, []int{2}},
		{1.0, []int{}, []int{1}},
		{1.5, []int{}, []int{1, 2}},
		{2.5, []int{1}, []int{2, 3, 4}},
		{3.5, []int{1, 5, 7}, []int{2, 3, 4}},
		{1.0, []int{1}, []int{1}},
		{1.5, []int{1, 2}, []int{1, 2}},
		{3.0, []int{1, 2}, []int{3, 4, 5}},
		{3.5, []int{1, 2}, []int{3, 4, 5, 6}},
	}

	for _, c := range cases {
		fmt.Println("===============", c.num1, c.num2)
		assert.Equal(t, c.answer, findMedianSortedArrays(c.num1, c.num2), "input: %v %v", c.num1, c.num2)
	}
}

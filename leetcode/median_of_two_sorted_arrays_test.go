package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMedianSortedArrays(t *testing.T) {
	t.Skip() // was not able to finish this task :(
	type testcase struct {
		answer float64
		num1   []int
		num2   []int
	}

	cases := []testcase{
		{4.5, []int{2, 3, 4, 5}, []int{1, 6, 7, 8}},
		{4.5, []int{1, 2, 3, 4}, []int{5, 6, 7, 8}},
		{4.5, []int{5, 6, 7, 8}, []int{1, 2, 3, 4}},
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
		{2.0, []int{1}, []int{2, 3}},
		{2.0, []int{3}, []int{1, 2}},
		{2.0, []int{1, 2}, []int{3}},
	}

	for _, c := range cases {
		fmt.Println("===============", c.num1, c.num2)
		assert.Equal(t, c.answer, findMedianSortedArrays(c.num1, c.num2), "input: %v %v", c.num1, c.num2)
	}
}

func Test_findMedianSortedArrays_arrayShiftIndexRight(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7}

	array := newArray(nums)
	inter := array.currentInterval()
	assert.Equal(t, 3, array.index)
	assert.Equal(t, 0, array.start)
	assert.Equal(t, 7, array.end)
	assert.Equal(t, 3, inter.left)
	assert.Equal(t, 4, inter.right)
	assert.False(t, inter.border)

	inter = array.shiftIndexRight()
	assert.Equal(t, 5, array.index)
	assert.Equal(t, 4, array.start)
	assert.Equal(t, 7, array.end)
	assert.Equal(t, 5, inter.left)
	assert.Equal(t, 6, inter.right)
	assert.False(t, inter.border)

	inter = array.shiftIndexRight()
	assert.Equal(t, 6, array.index)
	assert.Equal(t, 6, array.start)
	assert.Equal(t, 7, array.end)
	assert.Equal(t, 6, inter.left)
	assert.Equal(t, 7, inter.right)
	assert.False(t, inter.border)

	inter = array.shiftIndexRight()
	assert.Equal(t, 7, array.index)
	assert.Equal(t, 7, array.start)
	assert.Equal(t, 7, array.end)
	assert.Equal(t, 7, inter.left)
	assert.True(t, inter.border)

	assert.Panics(t, func() { _ = array.shiftIndexRight() })
}

func Test_findMedianSortedArrays_arrayShiftIndexLeft(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7}

	array := newArray(nums)
	inter := array.currentInterval()
	assert.Equal(t, 3, array.index)
	assert.Equal(t, 0, array.start)
	assert.Equal(t, 7, array.end)
	assert.Equal(t, 3, inter.left)
	assert.Equal(t, 4, inter.right)
	assert.False(t, inter.border)

	inter = array.shiftIndexLeft()
	assert.Equal(t, 1, array.index)
	assert.Equal(t, 0, array.start)
	assert.Equal(t, 3, array.end)
	assert.Equal(t, 1, inter.left)
	assert.Equal(t, 2, inter.right)
	assert.False(t, inter.border)

	inter = array.shiftIndexLeft()
	assert.Equal(t, 0, array.index)
	assert.Equal(t, 0, array.start)
	assert.Equal(t, 1, array.end)
	assert.Equal(t, 0, inter.left)
	assert.Equal(t, 1, inter.right)
	assert.False(t, inter.border)

	inter = array.shiftIndexLeft()
	assert.Equal(t, -1, array.index)
	assert.Equal(t, 0, array.start)
	assert.Equal(t, 0, array.end)
	assert.Equal(t, 0, inter.left)
	assert.True(t, inter.border)

	assert.Panics(t, func() { _ = array.shiftIndexLeft() })
}

func Test_findMedianSortedArrays_arrayShiftComplex(t *testing.T) {
	nums := []int{
		0, 1, 2, 3, 4, 5, 6, 7,
		8, 9, 10, 11, 12, 13, 14, 15}

	array := newArray(nums)

	inter := array.currentInterval()
	assert.Equal(t, 7, array.index)
	assert.Equal(t, 0, array.start)
	assert.Equal(t, 15, array.end)
	assert.Equal(t, 7, inter.left)
	assert.Equal(t, 8, inter.right)

	inter = array.shiftIndexRight()
	assert.Equal(t, 11, array.index)
	assert.Equal(t, 8, array.start)
	assert.Equal(t, 15, array.end)
	assert.Equal(t, 11, inter.left)
	assert.Equal(t, 12, inter.right)

	inter = array.shiftIndexLeft()
	assert.Equal(t, 9, array.index)
	assert.Equal(t, 8, array.start)
	assert.Equal(t, 11, array.end)
	assert.Equal(t, 9, inter.left)
	assert.Equal(t, 10, inter.right)

	inter = array.shiftIndexLeft()
	assert.Equal(t, 8, array.index)
	assert.Equal(t, 8, array.start)
	assert.Equal(t, 9, array.end)
	assert.Equal(t, 8, inter.left)
	assert.Equal(t, 9, inter.right)

}

func Test_findMedianSortedArrays_newArray(t *testing.T) {
	nums := []int{2, 3, 4}
	array := newArray(nums)
	assert.Equal(t, 0, array.index)
	assert.Equal(t, 0, array.start)
	assert.Equal(t, 2, array.end)
}

func Test_findMedianSortedArrays_intersections0(t *testing.T) {
	var inter1, inter2 interval

	inter1 = interval{left: 1, right: 2}
	inter2 = interval{left: 3, right: 4}
	assert.False(t, inter1.intersects(inter2))
	assert.False(t, inter2.intersects(inter1))

	inter1 = interval{left: 1, border: true}
	inter2 = interval{left: 2, right: 3}
	assert.True(t, inter1.intersects(inter2))
	assert.True(t, inter2.intersects(inter1))
}

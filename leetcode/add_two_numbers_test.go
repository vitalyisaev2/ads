package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_addTwoNumbersTest_basic(t *testing.T) {
	l1 := []*ListNode{
		{Val: 2},
		{Val: 4},
		{Val: 3},
	}
	l2 := []*ListNode{
		{Val: 5},
		{Val: 6},
		{Val: 4},
	}
	linkAdjacentNodes(l1)
	linkAdjacentNodes(l2)
	resultHead := addTwoNumbers(l1[0], l2[0])
	resultNodes := listToSlice(resultHead)

	assert.Len(t, resultNodes, 3)
	assert.Equal(t, resultNodes[0].Val, 7)
	assert.Equal(t, resultNodes[1].Val, 0)
	assert.Equal(t, resultNodes[2].Val, 8)
}

func Test_addTwoNumbersTest_lastPointOverflow1(t *testing.T) {
	l1 := []*ListNode{{Val: 5}}
	l2 := []*ListNode{{Val: 5}}
	resultHead := addTwoNumbers(l1[0], l2[0])
	resultNodes := listToSlice(resultHead)

	assert.Len(t, resultNodes, 2)
	assert.Equal(t, resultNodes[0].Val, 0)
	assert.Equal(t, resultNodes[1].Val, 1)
}

func Test_addTwoNumbersTest_lastPointOverflow2(t *testing.T) {
	l1 := []*ListNode{{Val: 9}}
	l2 := []*ListNode{{Val: 9}}
	resultHead := addTwoNumbers(l1[0], l2[0])
	resultNodes := listToSlice(resultHead)

	assert.Len(t, resultNodes, 2)
	assert.Equal(t, resultNodes[0].Val, 8)
	assert.Equal(t, resultNodes[1].Val, 1)
}

func linkAdjacentNodes(nodes []*ListNode) {
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Next = nodes[i+1]
	}
}

func listToSlice(head *ListNode) []*ListNode {
	var result []*ListNode
	for {
		if head == nil {
			break
		}
		result = append(result, head)
		head = head.Next
	}
	return result
}

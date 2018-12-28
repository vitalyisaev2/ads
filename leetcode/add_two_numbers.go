package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	var (
		overflow   bool
		head, next *ListNode
		op1, op2   = l1, l2
	)

	for {
		if op1 == nil && op2 == nil {
			if overflow {
				next.Next = &ListNode{Val: 1}
			}
			break
		}
		if head == nil {
			head, overflow = sumNodes(overflow, op1, op2)
			next = head
		} else {
			next.Next, overflow = sumNodes(overflow, op1, op2)
			next = next.Next
		}
		if op1 != nil {
			if op1.Next == nil {
				op1 = nil
			} else {
				op1 = op1.Next
			}
		}
		if op2 != nil {
			if op2.Next == nil {
				op2 = nil
			} else {
				op2 = op2.Next
			}
		}
	}
	return head
}

func sumNodes(overflowIn bool, nodes ...*ListNode) (lOut *ListNode, overflowOut bool) {
	result := 0
	for _, node := range nodes {
		if node != nil {
			result += node.Val
		}
	}
	if overflowIn {
		result++
	}
	if result > 9 {
		return &ListNode{Val: result % 10}, true
	}
	return &ListNode{Val: result}, false
}

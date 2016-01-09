// Skiena 3.2: stacks and queues

package structures

type StackElement interface{}

type Stack struct {
	LinkedList
}

// Push new element to the stack
func (stack *Stack) Push(item StackElement) {
	stack.Append(item)
}

// Pop last element from stack
func (stack *Stack) Pop() StackElement {

	// We need to reimplement deletion of the last element, since
	// LinkedList.Delete() is O(N).
	record := stack.last
	stack.last = record.next
	stack.length -= 1
	return record.item.(StackElement)
}

package problems

type StackMin struct {
	topElement *LinkedListNode
	minElement *LinkedListNode
}

func NewStackMin() *StackMin {
	return &StackMin{nil, nil}
}

func (stack *StackMin) Push(value int) {
	stack.topElement = NewLinkedListNode(value, stack.topElement)
	if stack.minElement == nil || stack.minElement.Data > value {
		stack.minElement = NewLinkedListNode(value, stack.minElement)
	}
}

func (stack *StackMin) Pop() (element int, isEmpty bool) {
	if stack.topElement == nil {
		return -1, true
	}

	if stack.topElement.Data == stack.minElement.Data {
		stack.minElement = stack.minElement.Next
	}

	element = stack.topElement.Data

	stack.topElement = stack.topElement.Next

	return element, false
}

func (stack *StackMin) Min() (int, bool) {
	if stack.minElement == nil {
		return -1, true
	}

	return stack.minElement.Data, false
}

func (stack *StackMin) TopElement() *LinkedListNode {
	return stack.topElement
}

func (stack *StackMin) MinElement() *LinkedListNode {
	return stack.minElement
}

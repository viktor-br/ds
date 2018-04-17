package problems

import (
	"github.com/viktor-br/ds/ds"
)

// SortStack with using another stack only
func SortStack(stack *ds.Stack) *ds.Stack {
	tempStack := ds.NewStack()

	var min1 int
	var min2 int
	elementsToProcess := stack.Length()

	if elementsToProcess <= 1 {
		return stack
	}

	var element int

	for elementsToProcess > 0 {
		// pour elements in temp stack
		element, _ = stack.Pop()
		min1 = element
		for i := 1; i < elementsToProcess; i++ {
			element, _ = stack.Pop()

			if min1 > element {
				tempStack.Push(min1)
				min1 = element
				continue
			}

			tempStack.Push(element)
		}
		elementsToProcess--

		// pour elements back
		element, _ = tempStack.Pop()
		min2 = element
		for i := 1; i < elementsToProcess; i++ {
			element, _ = tempStack.Pop()

			if min2 > element {
				stack.Push(min2)
				min2 = element
				continue
			}

			stack.Push(element)
		}
		elementsToProcess--

		tempStack.Push(min1)
		tempStack.Push(min2)

		if stack.Length() == 1 {
			element, _ = stack.Pop()
			tempStack.Push(element)
			elementsToProcess--
		}
	}

	for tempStack.Length() > 0 {
		element, _ = tempStack.Pop()
		stack.Push(element)
	}

	return stack
}

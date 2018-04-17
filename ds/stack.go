package ds

import (
	"fmt"
	"sync"
)

// Stack simple implementation
type Stack struct {
	data  []int
	mutex *sync.Mutex
}

// NewStack creates new stack
func NewStack() *Stack {
	return &Stack{data: make([]int, 0), mutex: &sync.Mutex{}}
}

// Push element to the stack
func (stack *Stack) Push(number int) {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	stack.data = append(stack.data, number)
}

// Pop top element of the stack
func (stack *Stack) Pop() (element int, isEmpty bool) {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	stackLength := len(stack.data)

	if stackLength == 0 {
		return -1, true
	}
	element = stack.data[stackLength-1]
	stack.data = stack.data[0 : stackLength-1]

	return element, false
}

// Length of the stack
func (stack *Stack) Length() int {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	return len(stack.data)
}

func (stack *Stack) Println() {
	fmt.Println(stack.data)
}

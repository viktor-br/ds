package main

import (
	"fmt"

	"github.com/viktor-br/ds/problems"
)

func main() {
	stack := problems.NewStackMin()

	stack.Push(3)
	stack.Push(4)
	stack.Push(2)
	stack.Push(1)

	fmt.Println(problems.LinkedListToString(stack.TopElement()))
	fmt.Println(problems.LinkedListToString(stack.MinElement()))

	fmt.Println(stack.Min())

	stack.Pop()
	fmt.Println(stack.Min())

	stack.Pop()
	fmt.Println(stack.Min())

	value, _ := stack.Pop()
	fmt.Println(stack.Min())
	fmt.Println(value)
}

package ds

import "testing"

func TestStack(t *testing.T) {
	stack := NewStack()
	stack.Push(5)
	stack.Push(6)
	stack.Push(7)

	if stack.Length() != 3 {
		t.Errorf("expect stack length 3")
	}

	if element, _ := stack.Pop(); element != 7 {
		t.Errorf("Expect 7 on top")
	}

	if element, _ := stack.Pop(); element != 6 {
		t.Errorf("Expect 6 on top")
	}

	if element, _ := stack.Pop(); element != 5 {
		t.Errorf("Expect 5 on top")
	}

	if _, isEmpty := stack.Pop(); !isEmpty {
		t.Errorf("Expect empty stack")
	}
}

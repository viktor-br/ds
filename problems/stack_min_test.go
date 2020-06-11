package problems

import "testing"

func TestStackMin1(t *testing.T) {
	stack := NewStackMin()

	stack.Push(3)
	stack.Push(4)
	stack.Push(2)
	stack.Push(1)

	stack.Pop()
	stack.Pop()
	// stack.Pop()

	actual, isEmpty := stack.Min()

	if actual != 3 {
		t.Errorf("expected: %d, actual: %d", 3, actual)
	}

	if isEmpty {
		t.Error("expected Min() isEmpty=false")
	}

	actual, _ = stack.Pop()
	if actual != 4 {
		t.Errorf("expected last value: 4, actual: %d", actual)
	}

	actual, _ = stack.Pop()
	if actual != 3 {
		t.Errorf("expected last value: 3, actual: %d", actual)
	}

}

func TestStackMinPopEmpty(t *testing.T) {
	stack := NewStackMin()

	_, isEmpty := stack.Pop()

	if !isEmpty {
		t.Error("expected pop() return isEmpty=true")
	}
}

func TestStackMinEmpty(t *testing.T) {
	stack := NewStackMin()

	_, isEmpty := stack.Min()

	if !isEmpty {
		t.Error("expected Min() return isEmpty=true")
	}
}

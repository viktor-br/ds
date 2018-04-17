package problems

import (
	"sort"
	"testing"

	"github.com/viktor-br/ds/ds"
)

func TestEmptyStackSort(t *testing.T) {
	stack := ds.NewStack()
	stack = SortStack(stack)
	if stack.Length() != 0 {
		t.Errorf("Magic happened")
	}
}

func TestOneElementStackSort(t *testing.T) {
	stack := ds.NewStack()
	stack.Push(1)
	stack = SortStack(stack)
	if stack.Length() != 1 {
		t.Errorf("Another magic happened")
	}
}

func TestOddElementsStackSort(t *testing.T) {
	inputData := []int{4, 2, 6, 8, 3, 9, 5}
	runTestStackSortForSlice(t, inputData)
}

func TestEvenElementsStackSort(t *testing.T) {
	inputData := []int{4, 2, 6, 8, 9, 5}
	runTestStackSortForSlice(t, inputData)
}

func runTestStackSortForSlice(t *testing.T, inputData []int) {
	expectedResult := make([]int, len(inputData))
	copy(expectedResult, inputData)
	sort.Ints(expectedResult)
	var element int
	stack := ds.NewStack()
	for i := 0; i < len(inputData); i++ {
		stack.Push(inputData[i])
	}

	if stack.Length() != len(inputData) {
		t.Errorf("expected length %d", len(inputData))
	}

	stack = SortStack(stack)
	for i := 0; i < stack.Length(); i++ {
		element, _ = stack.Pop()
		if element != expectedResult[i] {
			t.Errorf("stack[%d] = %d", i, element)
		}
	}
}

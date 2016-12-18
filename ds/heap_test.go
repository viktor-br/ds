package ds

import (
	"testing"
)

func TestHeapInsert(t *testing.T) {
	heap := NewHeap()

	for i := 31; i > 1; i-- {
		heap.Insert(i)
	}

	heap.Insert(1)

	heapData := heap.Data()
	heapLen := len(heapData)
	expectedHeapLen := 31
	if heapLen != expectedHeapLen {
		t.Errorf("Expect heap length %d, actual %d", expectedHeapLen, heapLen)
	}

	expectedFirstElement := 1
	if heapData[0] != expectedFirstElement {
		t.Errorf("Expect first element %d, actual %d", expectedFirstElement, heapData[0])
	}
}

func TestHeapDelete(t *testing.T) {
	heap := NewHeap()

	for i := 31; i > 0; i-- {
		heap.Insert(i)
	}

	heap.Delete(0)

	heapData := heap.Data()
	heapLen := len(heapData)
	expectedHeapLen := 30
	if heapLen != expectedHeapLen {
		t.Errorf("Expect heap length %d, actual %d", expectedHeapLen, heapLen)
	}

	expectedFirstElement := 2
	if heapData[0] != expectedFirstElement {
		t.Errorf("Expect first element %d, actual %d", expectedFirstElement, heapData[0])
	}
}

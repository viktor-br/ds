package ds

import (
	"errors"
	"sync"
)

// TODO implement generalized version of heap.

// Heap implement heap data structure for int type
type Heap struct {
	data []int
	mutex *sync.Mutex
}

// NewHeap creates and init new heap.
func NewHeap() *Heap {
	return &Heap{data: []int{}, mutex: &sync.Mutex{}}
}

// Data returns heap elements as []int
func (heap *Heap) Data() []int {
	heap.init()
	heap.mutex.Lock()
	defer heap.mutex.Unlock()

	return heap.data
}

// ParentOf element with index. -1 means no parent.
func (heap *Heap) ParentOf(index int) (int, error) {
	heap.init()
	heap.mutex.Lock()
	defer heap.mutex.Unlock()

	return heap.parentOf(index)
}

// MinChild returns index of minimum child for parentIndex. -1 means no min child.
func (heap *Heap) MinChild(parentIndex int) (int, error) {
	heap.init()
	heap.mutex.Lock()
	defer heap.mutex.Unlock()

	return heap.minChild(parentIndex)
}

// Insert element into heap.
func (heap *Heap) Insert(number int) error {
	heap.init()
	heap.mutex.Lock()
	defer heap.mutex.Unlock()

	lastIndex := len(heap.data)
	heap.data = append(heap.data, number)

	for {
		parentIndex, err := heap.parentOf(lastIndex)
		if err != nil {
			return err
		}
		if parentIndex == -1 {
			return nil
		}
		if heap.data[parentIndex] < number {
			return nil
		}
		heap.data[parentIndex], heap.data[lastIndex] = heap.data[lastIndex], heap.data[parentIndex]
		lastIndex = parentIndex
	}

	return nil
}

// Delete heap's element
func (heap *Heap) Delete(index int) error {
	heap.init()
	heap.mutex.Lock()
	defer heap.mutex.Unlock()

	l := len(heap.data)
	if l == 0 {
		return errors.New("Heap is empty")
	}
	if l == 1 {
		heap.data = []int{}
		return nil
	}
	if index >= l {
		return errors.New("Index is bigger then heap size")
	}

	heap.data[index] = heap.data[l - 1]
	heap.data = heap.data[0:l-1]
	for {
		childIndex, err := heap.minChild(index)
		if err != nil {
			return err
		}
		if childIndex == -1 {
			return nil
		}
		heap.data[index], heap.data[childIndex] = heap.data[childIndex], heap.data[index]
		index = childIndex
	}

	return nil
}

func (heap *Heap) init() {
	if heap.data == nil {
		heap.data = []int{}
	}
	if heap.mutex == nil {
		heap.mutex = &sync.Mutex{}
	}
}

func (heap *Heap) parentOf(index int) (int, error) {
	if index >= len(heap.data) {
		return 0, errors.New("Index is bigger then heap size")
	}
	if index == 0 {
		return -1, nil
	}
	index = (index + 1) >> 1

	return index - 1, nil
}

func (heap *Heap) minChild(parentIndex int) (int, error) {
	heapLen := len(heap.data)
	if parentIndex >= heapLen {
		return 0, errors.New("Index is bigger then heap size")
	}

	child1 := (parentIndex << 1) + 1
	child2 := child1 + 1

	if child1 >= heapLen {
		return -1, nil
	}

	if child2 >= heapLen {
		if heap.data[parentIndex] < heap.data[child1] {
			return -1, nil
		}
		return child1, nil
	}
	if heap.data[parentIndex] < heap.data[child1] && heap.data[parentIndex] < heap.data[child2] {
		return -1, nil
	}
	if heap.data[child1] > heap.data[child2] {
		return child2, nil
	}

	return child1, nil
}
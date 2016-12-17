package sort

import (
	"testing"
	"sort"
)

func benchmarkQuickSort(a []int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort(a)
	}
}

func TestQuickSort(t *testing.T) {
	a := []int{1, 11, 7, 3, 6, 2, 8, 4, 9, 10, 5, 7, 1}
	b := make([]int, len(a))
	copy(b, a)
	sort.Ints(b[:])

	QuickSort(a)
	if !Compare(a, b) {
		t.Errorf("Array is not sorted properly. Expected: %v, actual: %v", b, a)
	}

}

func BenchmarkQuickSort1000(b *testing.B) {
	a := generateArray(1000)
	benchmarkQuickSort(a, b)
}

func BenchmarkQuickSort10000(b *testing.B) {
	a := generateArray(10000)
	benchmarkQuickSort(a, b)
}

func BenchmarkQuickSort100000(b *testing.B) {
	a := generateArray(100000)
	benchmarkQuickSort(a, b)
}

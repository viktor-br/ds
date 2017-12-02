package sort

import (
	"sort"
	"testing"
)

func benchmarkMergeSort(a []int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSort(a)
	}
}

func TestMergeSort(t *testing.T) {
	a := []int{1, 11, 7, 3, 6, 2, 8, 4, 9, 6, 10, 5, 1}
	b := make([]int, len(a))
	copy(b, a)
	sort.Ints(b)

	result := MergeSort(a)
	if !Compare(result, b) {
		t.Errorf("Array is not sorted properly. Expected: %v, actual: %v", b, result)
	}
}

func BenchmarkMergeSort1000(b *testing.B) {
	a := generateArray(1000)
	benchmarkMergeSort(a, b)
}

func BenchmarkMergeSort10000(b *testing.B) {
	a := generateArray(10000)
	benchmarkMergeSort(a, b)
}

func BenchmarkMergeSort100000(b *testing.B) {
	a := generateArray(100000)
	benchmarkMergeSort(a, b)
}

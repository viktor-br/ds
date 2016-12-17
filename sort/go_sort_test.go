package sort

import (
	"testing"
	"sort"
)

func benchmarkGoSort(a []int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.Ints(a)
	}
}

func BenchmarkGoSort1000(b *testing.B) {
	a := generateArray(1000)
	benchmarkGoSort(a, b)
}

func BenchmarkGoSort10000(b *testing.B) {
	a := generateArray(10000)
	benchmarkGoSort(a, b)
}

func BenchmarkGoSort100000(b *testing.B) {
	a := generateArray(100000)
	benchmarkGoSort(a, b)
}


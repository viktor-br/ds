package sort

import "math/rand"

func generateArray(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = rand.Intn(n)
	}

	return a
}

func Compare(a, b []int) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

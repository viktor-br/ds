package problems

import (
	"math/rand"
	"testing"
	"time"
)

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func generateArray(n int) []int {
	var k int
	var j int
	a := make([]int, n)

	for i := 1; i < n; i++ {
		a[i] = i + 1
	}
	for i := 0; i < n; i++ {
		k = rnd.Intn(n)
		j = rnd.Intn(n)
		a[k], a[j] = a[j], a[k]
	}

	return a
}

func BenchmarkBestStocks1000(b *testing.B) {
	a := generateArray(1000)
	for i := 0; i < b.N; i++ {
		BestStocks(a)
	}
}

func BenchmarkBestStocks10000(b *testing.B) {
	a := generateArray(10000)
	for i := 0; i < b.N; i++ {
		BestStocks(a)
	}
}

func BenchmarkBestStocks100000(b *testing.B) {
	a := generateArray(100000)
	for i := 0; i < b.N; i++ {
		BestStocks(a)
	}
}

func BenchmarkBestStocksBruteForce1000(b *testing.B) {
	a := generateArray(1000)
	for i := 0; i < b.N; i++ {
		BestStocksBruteForce(a)
	}
}

func BenchmarkBestStocksBruteForce10000(b *testing.B) {
	a := generateArray(10000)
	for i := 0; i < b.N; i++ {
		BestStocksBruteForce(a)
	}
}

func BenchmarkBestStocksBruteForce100000(b *testing.B) {
	a := generateArray(100000)
	for i := 0; i < b.N; i++ {
		BestStocksBruteForce(a)
	}
}

func TestBestStocks(t *testing.T) {
	a := []int{3, 4, 6, 7, 9, 5, 8, 2, 1, 0}

	expectedI, expectedJ, expectedAmount := 0, 4, 6

	i, j := BestStocks(a)
	amount := a[j] - a[i]

	if i != expectedI {
		t.Errorf("Expected i: %d, actual i: %d", expectedI, i)
	}

	if j != expectedJ {
		t.Errorf("Expected j: %d, actual j: %d", expectedJ, j)
	}

	if amount != expectedAmount {
		t.Errorf("Expected amount: %d, actual amount: %d", expectedAmount, amount)
	}
}

func TestBestStocksBruteForce(t *testing.T) {
	a := []int{3, 4, 6, 7, 9, 5, 8, 2, 1, 0}

	expectedI, expectedJ, expectedAmount := 0, 4, 6

	i, j := BestStocksBruteForce(a)
	amount := a[j] - a[i]

	if i != expectedI {
		t.Errorf("Expected i: %d, actual i: %d", expectedI, i)
	}

	if j != expectedJ {
		t.Errorf("Expected j: %d, actual j: %d", expectedJ, j)
	}

	if amount != expectedAmount {
		t.Errorf("Expected amount: %d, actual amount: %d", expectedAmount, amount)
	}
}

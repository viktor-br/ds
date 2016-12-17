package problems

import (
	"testing"
	"fmt"
)

//var rnd *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func TestMultOthers(t *testing.T) {
	multTotal := 1;
	n := 10000
	a := make([]int, n)

	for i := 0; i < n; i++ {
		a[i] = rnd.Intn(10)
		multTotal = multTotal * a[i]
	}

	indx := rnd.Intn(n)
	multExpected := multTotal / a[indx]

	multActual := MultOthers(a, indx)

	if multExpected != multActual {
		fmt.Errorf("Expected %d, actual %d", multExpected, multActual)
	}
}

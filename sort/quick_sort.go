package sort

import (
	"math/rand"
)
import "time"

var rnd *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func QuickSort(a []int) {
	aLen := len(a)

	if aLen <= 1 {
		return
	}

	pivotIndex := rnd.Intn(aLen)
	pivotValue := a[pivotIndex]

	i := 1 //  < pivot border
	j := 1 // > pivot border

	a[0], a[pivotIndex] = a[pivotIndex], a[0]

	for k := j; k < aLen; k++ {
		if a[k] <= pivotValue {
			if a[k] != a[i] {
				a[k], a[i] = a[i], a[k]
			}
			i++
		}
		j++
	}
	if a[i-1] != a[0] {
		a[i-1], a[0] = a[0], a[i-1]
	}

	QuickSort(a[0 : i-1])
	QuickSort(a[i:])
}

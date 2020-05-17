package problems

import (
	"testing"

	"github.com/viktor-br/ds/util"
)

func TestZeroMatrix(t *testing.T) {
	n := 4
	matrix := make([][]int, n)
	matrix[0] = []int{1, 1, 1, 1}
	matrix[1] = []int{1, 0, 1, 1}
	matrix[2] = []int{1, 1, 1, 1}
	matrix[3] = []int{1, 1, 0, 1}

	expectedMatrix := make([][]int, n)
	expectedMatrix[0] = []int{1, 0, 0, 1}
	expectedMatrix[1] = []int{0, 0, 0, 0}
	expectedMatrix[2] = []int{1, 0, 0, 1}
	expectedMatrix[3] = []int{0, 0, 0, 0}

	testCases := []struct {
		Name     string
		Matrix   [][]int
		Expected [][]int
	}{
		{"zero matrix", matrix, expectedMatrix},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			actualMatrix := ZeroMatrix(tc.Matrix)
			if !util.EqualMatrices(actualMatrix, expectedMatrix) {
				t.Errorf("Expect \n%s \n Actual \n%s \n", util.MatrixToString(expectedMatrix), util.MatrixToString(actualMatrix))
			}
		})
	}
}

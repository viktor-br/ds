package problems

import "testing"

func TestMatrixRotateOdd(t *testing.T) {
	matrix := make([][]int, 3)
	matrix[0] = []int{1, 2, 3}
	matrix[1] = []int{4, 5, 6}
	matrix[2] = []int{7, 8, 9}

	expected := make([][]int, 3)
	expected[0] = []int{7, 4, 1}
	expected[1] = []int{8, 5, 2}
	expected[2] = []int{9, 6, 3}

	runMatrixRotateTest(t, matrix, expected)
}

func TestMatrixRotateEven(t *testing.T) {
	matrix := make([][]int, 4)
	matrix[0] = []int{11, 12, 13, 14}
	matrix[1] = []int{15, 16, 17, 18}
	matrix[2] = []int{19, 20, 21, 22}
	matrix[3] = []int{23, 24, 25, 26}

	expected := make([][]int, 4)
	expected[0] = []int{23, 19, 15, 11}
	expected[1] = []int{24, 20, 16, 12}
	expected[2] = []int{25, 21, 17, 13}
	expected[3] = []int{26, 22, 18, 14}

	runMatrixRotateTest(t, matrix, expected)
}

func runMatrixRotateTest(t *testing.T, matrix [][]int, expected [][]int) {
	actual, err := RotateMatrix(matrix)
	if err != nil {
		t.Errorf("Matrix is not square")
	}

	if !equalMatrices(actual, expected) {
		t.Errorf("Expected is not actual for matrix rotation, n=3")
	}

	actual, err = RotateMatrixInPlace(matrix)
	if err != nil {
		t.Errorf("Matrix is not square")
	}
	if !equalMatrices(actual, expected) {
		t.Errorf("Expected is not actual for the in place matrix rotation, n=3")
	}
}

func equalMatrices(m1 [][]int, m2 [][]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	n := len(m1)
	for i := 0; i < len(m1); i++ {
		if len(m1[i]) != len(m2[i]) {
			return false
		}
		if len(m1[i]) != n {
			return false
		}
		for j := 0; j < len(m1[i]); j++ {
			if m1[i][j] != m2[i][j] {
				return false
			}
		}
	}

	return true
}

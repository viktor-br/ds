package problems

import (
	"fmt"
	"math"
)

// RotateMatrix rotate matrix by creating new one, uses twice more space.
func RotateMatrix(matrix [][]int) ([][]int, error) {
	n := len(matrix)
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}
	var x, y int
	dimensionEven := n%2 == 0
	center := int(math.Floor(float64(n) / 2))
	for i := 0; i < n; i++ {
		if len(matrix[i]) != n {
			return nil, fmt.Errorf("Expect nxn matrix")
		}
		for j := 0; j < n; j++ {
			x, y = i, j
			x, y = RotatePoint(x, y, center, dimensionEven)
			// fmt.Printf("%d:%d;%d:%d\n", i, j, x, y)
			result[x][y] = matrix[i][j]
		}
	}

	return result, nil
}

// RotateMatrixInPlace rotate matrix without additional matrix or variable
func RotateMatrixInPlace(matrix [][]int) ([][]int, error) {
	var x, y, j int
	n := len(matrix)
	m := n - 1
	center := int(math.Floor(float64(n) / 2))
	dimensionEven := n%2 == 0
	for i := 0; i < center; i++ {
		for z := 0; z < m; z++ {
			x = i
			y = i + z
			j = i + z

			for k := 0; k < 3; k++ {
				x, y = RotatePoint(x, y, center, dimensionEven)
				// fmt.Printf("(i,j):(%d,%d)=%d; (x,y):(%d,%d)=%d;\n\n", i, j, matrix[i][j], x, y, matrix[x][y])
				if len(matrix[i]) != n || len(matrix[x]) != n {
					return nil, fmt.Errorf("Expect nxn matrix")
				}
				matrix[i][j], matrix[x][y] = matrix[x][y], matrix[i][j]
			}
		}
		m = m - 2
	}

	return matrix, nil
}

// RotatePoint returns x, y coordinates for 90 degree clockwise rotation for one point.
func RotatePoint(x, y, center int, dimensionEven bool) (int, int) {
	if x >= center && dimensionEven {
		x = x + 1
	}
	if y >= center && dimensionEven {
		y = y + 1
	}
	x, y = y, -x+2*center
	if x >= center && dimensionEven {
		x = x - 1
	}
	if y >= center && dimensionEven {
		y = y - 1
	}

	return x, y
}

package problems

// RotateMatrixInPlace rotate matrix without additional matrix or variable
func RotateMatrixInPlace(matrix [][]int) ([][]int, error) {
	var x, y, j int
	n := len(matrix)
	m := n - 1
	center := n / 2
	for i := 0; i < center; i++ {
		for z := 0; z < m; z++ {
			j = i + z
			x, y = i+z, i+m
			matrix[i][j], matrix[x][y] = matrix[x][y], matrix[i][j]

			x, y = i+m, y-z
			matrix[i][j], matrix[x][y] = matrix[x][y], matrix[i][j]

			x, y = i+m-z, i
			matrix[i][j], matrix[x][y] = matrix[x][y], matrix[i][j]
		}
		m = m - 2
	}

	return matrix, nil
}

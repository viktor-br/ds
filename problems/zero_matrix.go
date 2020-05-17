package problems

func ZeroMatrix(matrix [][]int) [][]int {
	var zeroRows []int
	var zeroColumns []int
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				zeroRows = append(zeroRows, i)
				zeroColumns = append(zeroColumns, j)
			}
		}
	}
	for k := 0; k < len(zeroRows); k++ {
		for j := 0; j < len(matrix[zeroRows[k]]); j++ {
			matrix[zeroRows[k]][j] = 0
		}
	}
	for k := 0; k < len(zeroColumns); k++ {
		for i := 0; i < len(matrix[zeroColumns[k]]); i++ {
			matrix[i][zeroColumns[k]] = 0
		}
	}

	return matrix
}

package util

import (
	"fmt"
	"strings"
)

func EqualMatrices(m1 [][]int, m2 [][]int) bool {
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

func MatrixToString(matrix [][]int) string {
	var builder strings.Builder
	builder.Grow(32)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Fprintf(&builder, "%d ", matrix[i][j])
		}
		fmt.Fprintf(&builder, "\n")
	}

	return builder.String()
}

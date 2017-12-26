package problems

import (
	"math"
)

// SingleTransformation checks if one string is a single-transformation of another --
// it could be created by one (zero) action of insert, delete or replace
func SingleTransformation(first string, second string) bool {
	firstLen := len(first)
	secondLen := len(second)
	missedLetters := 0
	if math.Abs(float64(firstLen-secondLen)) > 1 {
		return false
	}

	i := 0
	j := 0
	for i < firstLen && j < secondLen && missedLetters < 2 {
		// fmt.Printf("%d:%d\n", i, j)
		if first[i] != second[j] {
			missedLetters = missedLetters + 1
			if firstLen == secondLen {
				i = i + 1
				j = j + 1
			} else if firstLen > secondLen {
				i = i + 1
			} else {
				j = j + 1
			}
		} else {
			i = i + 1
			j = j + 1
		}
	}

	return missedLetters < 2
}

package problems

import "strings"

func IsStringRotate(initial, rotated string) bool {
	s := rotated + rotated
	return strings.Contains(s, initial)
}

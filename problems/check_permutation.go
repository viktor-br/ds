package problems

import "unicode/utf8"

func CheckPermutation(s1, s2 string) bool {
	s1Len := len(s1)
	s2Len := len(s2)
	if s1Len != s2Len {
		return false
	}
	lettersCounter := make(map[byte]int)
	for i := 0; i < s1Len; i++ {
		lettersCounter[s1[i]]++
	}

	for i := 0; i < s2Len; i++ {
		lettersCounter[s2[i]]--
		if lettersCounter[s2[i]] < 0 {
			return false
		}
	}

	return true
}

func CheckPermutationUtf8(s1, s2 string) bool {
	s1Len := utf8.RuneCountInString(s1)
	s2Len := utf8.RuneCountInString(s2)
	if s1Len != s2Len {
		return false
	}
	lettersCounter := make(map[rune]int)
	for _, currentRune := range s1 {
		lettersCounter[currentRune]++
	}

	for _, currentRune := range s2 {
		lettersCounter[currentRune]--
		if lettersCounter[currentRune] < 0 {
			return false
		}
	}

	return true
}

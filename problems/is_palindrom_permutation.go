package problems

// import "fmt"

func IsPalindromPermutation(s string) bool {
	letterOddEven := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		letterOddEven[s[i]] = !letterOddEven[s[i]]
	}
	hasOddLetterNumber := false
	for _, v := range letterOddEven {
		if !v {
			continue
		}
		if hasOddLetterNumber {
			return false
		}
		hasOddLetterNumber = true
	}

	return true
}

func IsPalindromPermutationUtf8(s string) bool {
	letterOddEven := make(map[rune]bool)
	for _, currentRune := range s {
		letterOddEven[currentRune] = !letterOddEven[currentRune]
	}
	hasOddLetterNumber := false
	for _, v := range letterOddEven {
		if !v {
			continue
		}
		if hasOddLetterNumber {
			return false
		}
		hasOddLetterNumber = true
	}

	return true
}

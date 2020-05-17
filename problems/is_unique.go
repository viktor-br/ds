package problems

func IsUnique(s string) bool {
	strLen := len(s)
	symbols := make(map[byte]bool)

	for i := 0; i < strLen; i++ {
		if symbols[s[i]] {
			return false
		}

		symbols[s[i]] = true
	}

	return true
}

func IsUniqueUtf8(s string) bool {
	symbols := make(map[rune]bool)
	for _, currentRune := range s {
		_, ok := symbols[currentRune]
		if ok {
			return false
		}

		symbols[currentRune] = true
	}

	return true
}

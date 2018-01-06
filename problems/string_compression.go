package problems

import (
	"fmt"
	"strconv"
)

// StringCompression simple string compression aaabbc => a3b2c only for ASCII
func StringCompression(s string) string {
	result := []byte{}
	count := 1
	letter := s[0]
	fmt.Println(s)
	for i := 1; i < len(s); i++ {
		// fmt.Printf("%s:%s:%d\n", string(s[i-1]), string(s[i]), count)
		if s[i-1] == s[i] {
			count = count + 1
			continue
		}

		result = append(result, letter)
		if count > 1 {
			result = append(result, strconv.Itoa(count)...)
		}
		count = 1
		letter = s[i]
	}
	result = append(result, letter)
	if count > 1 {
		result = append(result, strconv.Itoa(count)...)
	}

	return string(result)
}

// StringCompressionUtf8 simple string compression aaabbc => a3b2c for utf8
func StringCompressionUtf8(s string) string {
	if len(s) == 0 {
		return s
	}
	var prevRune rune
	result := []rune{}
	count := 1
	for i, currentRune := range s {
		if i == 0 {
			prevRune = currentRune
			continue
		}
		// fmt.Printf("%s:%s:%d\n", string(s[i-1]), string(s[i]), count)
		if prevRune == currentRune {
			count = count + 1
			continue
		}

		result = append(result, prevRune)
		if count > 1 {
			result = append(result, []rune(strconv.Itoa(count))...)
		}
		count = 1
		prevRune = currentRune
	}
	result = append(result, prevRune)
	if count > 1 {
		result = append(result, []rune(strconv.Itoa(count))...)
	}

	return string(result)
}

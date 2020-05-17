package problems

func Urlify(s []byte, l int) []byte {
	current := len(s) - 1
	for i := l - 1; i >= 0; i-- {
		if s[i] == 32 {
			s[current] = 48
			current--
			s[current] = 50
			current--
			s[current] = 37 // %
			current = current - 1
		} else {
			s[current] = s[i]
			current--
		}
	}

	return s
}

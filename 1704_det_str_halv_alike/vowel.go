package vowel

func halvesAreAlike(s string) bool {
	var a, b int
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if isVovel(s[i]) {
			a++
		}
		if isVovel(s[j]) {
			b++
		}
	}
	return (a == b)
}

func isVovel(c byte) bool {
	switch c {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	case 'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}

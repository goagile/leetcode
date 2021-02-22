package roman

// ToInt - parse Roman dogit to Integer digit
func ToInt(s string) int {
	z := 0
	i, n := 0, len(s)
	a, b := s[0], byte(0)

	if n == 1 {
		return parse(a)
	}

	for i < (n - 1) {
		a, b = s[i], s[i+1]
		if pair(a, b) {
			z += parse(b) - parse(a)
			i += 2
			continue
		}
		z += parse(a)
		i++
	}

	if i < n {
		z += parse(s[n-1])
	}

	return z
}

func parse(c byte) int {
	switch c {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	}
	return 0
}

func pair(a, b byte) bool {
	switch {
	case a == 'I' && (b == 'V' || b == 'X'):
		return true
	case a == 'X' && (b == 'L' || b == 'C'):
		return true
	case a == 'C' && (b == 'D' || b == 'M'):
		return true
	}
	return false
}

package split

func balancedStringSplit(s string) int {
	var L, R, c int
	for _, b := range s {
		if b == 'R' {
			R++
		} else if b == 'L' {
			L++
		}
		if R == L {
			c++
		}
	}
	return c
}

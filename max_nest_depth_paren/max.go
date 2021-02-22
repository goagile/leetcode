package max

func maxDepth(s string) int {
	var m, k int
	for _, c := range s {
		switch c {
		case '(':
			k++
			if k > m {
				m = k
			}
		case ')':
			k--
		}
	}
	return m
}

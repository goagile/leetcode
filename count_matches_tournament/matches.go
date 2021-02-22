package matches

func numberOfMatches(n int) int {
	var matches, total int
	for n > 1 {
		if 0 == n%2 {
			matches = n / 2
			n = matches
		} else {
			matches = (n - 1) / 2
			n = matches + 1
		}
		total += matches
	}
	return total
}

func numberOfMatchesR(n int) int {
	var total int
	even := func(b int) bool { return 0 == n%2 }
	matches := func(n int) int {
		if even(n) {
			return n / 2
		}
		return (n - 1) / 2
	}
	for n > 1 {
		m := matches(n)
		total += m
		if even(n) {
			n = m
		} else {
			n = m + 1
		}
	}
	return total
}

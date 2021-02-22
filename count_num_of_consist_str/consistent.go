package consistent

func countConsistentStrings(allowed string, words []string) int {
	set := map[rune]bool{}
	for _, c := range allowed {
		set[c] = true
	}
	onlyallowed := func(word string) bool {
		for _, c := range word {
			if _, ok := set[c]; !ok {
				return false
			}
		}
		return true
	}
	c := 0
	for _, word := range words {
		if onlyallowed(word) {
			c++
		}
	}
	return c
}

func countConsistentStringsAlpha(allowed string, words []string) int {
	set := make([]int, 26, 26)
	for _, c := range allowed {
		set[c-'a'] = 1
	}
	onlyallowed := func(word string) bool {
		for _, c := range word {
			if set[c-'a'] == 0 {
				return false
			}
		}
		return true
	}
	c := 0
	for _, word := range words {
		if onlyallowed(word) {
			c++
		}
	}
	return c
}

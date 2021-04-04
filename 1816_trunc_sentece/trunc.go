package trunc

import "strings"

func truncateSentence(s string, k int) string {
	w := 0
	b := new(strings.Builder)
	for _, c := range s {
		if c == ' ' {
			w++
		}
		if w >= k {
			break
		}
		b.WriteRune(c)
	}
	return b.String()
}

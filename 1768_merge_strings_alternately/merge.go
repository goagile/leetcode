package merge

import "strings"

func mergeAlternately(word1 string, word2 string) string {
	b := new(strings.Builder)
	i := 0
	for ; i < len(word1) && i < len(word2); i++ {
		b.WriteByte(word1[i])
		b.WriteByte(word2[i])
	}
	if len(word1) > len(word2) {
		for ; i < len(word1); i++ {
			b.WriteByte(word1[i])
		}
	} else {
		for ; i < len(word2); i++ {
			b.WriteByte(word2[i])
		}
	}
	return b.String()
}

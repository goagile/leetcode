package check

import "strings"

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	w1s := ""
	for _, s := range word1 {
		w1s += s
	}
	w2s := ""
	for _, s := range word2 {
		w2s += s
	}
	return w1s == w2s
}

func arrayStringsAreEqualJoin(w1 []string, w2 []string) bool {
	return strings.Join(w1, "") == strings.Join(w2, "")
}

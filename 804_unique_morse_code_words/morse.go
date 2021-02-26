package morse

import "strings"

var table = []string{
	".-", "-...", "-.-.", "-..", ".",
	"..-.", "--.", "....", "..", ".---",
	"-.-", ".-..", "--", "-.", "---",
	".--.", "--.-", ".-.", "...", "-",
	"..-", "...-", ".--", "-..-", "-.--",
	"--..",
}

func uniqueMorseRepresentations(words []string) int {
	unique := 0
	seen := make(map[string]bool)
	for _, word := range words {
		b := new(strings.Builder)
		for _, c := range word {
			b.WriteString(table[c-'a'])
		}
		encoded := b.String()
		if _, ok := seen[encoded]; !ok {
			unique++
			seen[encoded] = true
		}
	}
	return unique
}

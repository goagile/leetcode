package alph

import (
	"strconv"
	"strings"
)

func freqAlphabets(s string) string {
	b := new(strings.Builder)
	step := 1
	for i := 0; i < len(s); i += step {
		x := string(s[i])
		if (i+2) < len(s) && s[i+2] == '#' {
			x += string(s[i+1])
			step = 3
		} else {
			step = 1
		}
		z, _ := strconv.Atoi(x)
		b.WriteRune(rune('a') + rune(z) - 1)
	}
	return b.String()
}

package tolowcase

import (
	"bytes"
	"strings"
	"unicode"
)

const shift = 'a' - 'A'

func toLowerCaseStringsBuilder(s string) string {
	b := new(strings.Builder)
	b.Grow(len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		if 'A' <= c && c <= 'Z' {
			c += shift
		}
		b.WriteByte(c)
	}
	return b.String()
}

func toLowerCaseStd(str string) string {
	return strings.ToLower(str)
}

func toLowerCaseBytesBuffer(str string) string {
	var b bytes.Buffer
	for _, r := range str {
		b.WriteRune(unicode.ToLower(r))
	}
	return b.String()
}

package shuffle

import "sort"

func restoreString(s string, indices []int) string {
	pairs := []Pair{}
	for i, pos := range indices {
		pairs = append(pairs, Pair{pos: pos, run: s[i]})
	}
	sort.Sort(ByPos(pairs))
	z := []byte{}
	for _, pair := range pairs {
		z = append(z, pair.run)
	}
	return string(z)
}

type Pair struct {
	pos int
	run byte
}

type ByPos []Pair

func (b ByPos) Len() int {
	return len(b)
}

func (b ByPos) Less(i, j int) bool {
	return b[i].pos < b[j].pos
}

func (b ByPos) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func restoreStringBytes(s string, indices []int) string {
	z := make([]byte, len(s))
	for i, v := range indices {
		z[v] = s[i]
	}
	return string(z)
}

func restoreStringRunes(s string, indices []int) string {
	z := make([]rune, len(s))
	for i, c := range s {
		z[indices[i]] = c
	}
	return string(z)
}

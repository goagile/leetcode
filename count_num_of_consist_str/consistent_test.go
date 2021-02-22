package consistent

import (
	"testing"
)

var examples = []struct {
	want    int
	allowed string
	words   []string
}{
	{
		want:    2,
		allowed: "ab",
		words:   []string{"ad", "bd", "aaab", "baa", "badab"},
	},
	{
		want:    7,
		allowed: "abc",
		words:   []string{"a", "b", "c", "ab", "ac", "bc", "abc"},
	},
	{
		want:    4,
		allowed: "cad",
		words:   []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"},
	},
}

var bench = examples[0]

func Test_examples_countConsistentStrings(t *testing.T) {
	for i, e := range examples {
		got := countConsistentStrings(e.allowed, e.words)
		if got != e.want {
			t.Fatalf("\ni:%v\nwant%v\ngot: %v\n", i, e.want, got)
		}
	}
}

func Benchmark_examples_countConsistentStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countConsistentStrings(bench.allowed, bench.words)
	}
}

func Test_examples_countConsistentStringsAlpha(t *testing.T) {
	for i, e := range examples {
		got := countConsistentStringsAlpha(e.allowed, e.words)
		if got != e.want {
			t.Fatalf("\ni:%v\nwant%v\ngot: %v\n", i, e.want, got)
		}
	}
}

func Benchmark_examples_countConsistentStringsAlpha(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countConsistentStringsAlpha(bench.allowed, bench.words)
	}
}

package merge

import "testing"

var examples = []struct {
	want  string
	word1 string
	word2 string
}{
	{
		want:  "apbqcr",
		word1: "abc",
		word2: "pqr",
	},
	{
		want:  "apbqrs",
		word1: "ab",
		word2: "pqrs",
	},
	{
		want:  "apbqcd",
		word1: "abcd",
		word2: "pq",
	},
	{
		want:  "ab",
		word1: "a",
		word2: "b",
	},
	{
		want:  "",
		word1: "",
		word2: "",
	},
}

func Test_examples_mergeAlternately(t *testing.T) {
	for i, e := range examples {
		got := mergeAlternately(e.word1, e.word2)
		if got != e.want {
			t.Fatalf(
				"\ni:%v\ngot:%v\nwant:%v\n",
				i, got, e.want,
			)
		}
	}
}

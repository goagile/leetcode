package check

import (
	"reflect"
	"testing"
)

var examples = []struct {
	want         bool
	word1, word2 []string
}{
	{
		want:  true,
		word1: []string{"ab", "c"},
		word2: []string{"a", "bc"},
	},
	{
		want:  false,
		word1: []string{"a", "cb"},
		word2: []string{"ab", "c"},
	},
	{
		want:  true,
		word1: []string{"abc", "d", "defg"},
		word2: []string{"abcddefg"},
	},
}

var bench = examples[0]

func Test_examples(t *testing.T) {
	for i, e := range examples {
		got := arrayStringsAreEqual(e.word1, e.word2)
		if !reflect.DeepEqual(e.want, got) {
			t.Fatalf("\ni:%v\nwant%v\ngot: %v\n", i, e.want, got)
		}
	}
}

func Benchmark_arrayStringsAreEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arrayStringsAreEqual(bench.word1, bench.word2)
	}
}

func Test_examples_arrayStringsAreEqualJoin(t *testing.T) {
	for i, e := range examples {
		got := arrayStringsAreEqualJoin(e.word1, e.word2)
		if !reflect.DeepEqual(e.want, got) {
			t.Fatalf("\ni:%v\nwant%v\ngot: %v\n", i, e.want, got)
		}
	}
}

func Benchmark_arrayStringsAreEqualJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arrayStringsAreEqualJoin(bench.word1, bench.word2)
	}
}

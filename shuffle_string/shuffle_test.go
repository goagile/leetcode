package shuffle

import "testing"

var examples = []struct {
	want    string
	indices []int
	s       string
}{
	{
		want:    "leetcode",
		indices: []int{4, 5, 6, 7, 0, 2, 1, 3},
		s:       "codeleet",
	},
	{
		want:    "abc",
		indices: []int{0, 1, 2},
		s:       "abc",
	},
	{
		want:    "nihao",
		indices: []int{3, 1, 4, 2, 0},
		s:       "aiohn",
	},
	{
		want:    "arigatou",
		indices: []int{4, 0, 2, 6, 7, 3, 1, 5},
		s:       "aaiougrt",
	},
	{
		want:    "rat",
		indices: []int{1, 0, 2},
		s:       "art",
	},
}

var bench = examples[0]

func Test_examples_restoreString(t *testing.T) {
	for i, e := range examples {
		got := restoreString(e.s, e.indices)
		if got != e.want {
			t.Fatalf("\ni:%vgot :%v\nwant:%v\n", i, got, e.want)
		}
	}
}

func Benchmark_restoreString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		restoreString(bench.s, bench.indices)
	}
}

func Test_examples_restoreStringBytes(t *testing.T) {
	for i, e := range examples {
		got := restoreStringBytes(e.s, e.indices)
		if got != e.want {
			t.Fatalf("\ni:%vgot :%v\nwant:%v\n", i, got, e.want)
		}
	}
}

func Benchmark_restoreStringBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		restoreStringBytes(bench.s, bench.indices)
	}
}

func Test_examples_restoreStringRunes(t *testing.T) {
	for i, e := range examples {
		got := restoreStringRunes(e.s, e.indices)
		if got != e.want {
			t.Fatalf("\ni:%vgot :%v\nwant:%v\n", i, got, e.want)
		}
	}
}

func Benchmark_restoreStringRunes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		restoreStringRunes(bench.s, bench.indices)
	}
}

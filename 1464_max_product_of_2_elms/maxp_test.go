package maxp

import "testing"

var examples = []struct {
	want int
	nums []int
}{
	{
		want: 12,
		nums: []int{3, 4, 5, 2},
	},
	{
		want: 16,
		nums: []int{1, 5, 4, 5},
	},
	{
		want: 12,
		nums: []int{3, 7},
	},
}

var bench = examples[0]

func Test_examples_maxProduct(t *testing.T) {
	for i, e := range examples {
		got := maxProduct(e.nums)
		if got != e.want {
			t.Fatalf(
				"\ni:%vgot:%v\nwant:%v\n",
				i, got, e.want,
			)
		}
	}
}

func Benchmark_examples_maxProduct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxProduct(bench.nums)
	}
}

func Test_examples_maxProductSort(t *testing.T) {
	for i, e := range examples {
		got := maxProductSort(e.nums)
		if got != e.want {
			t.Fatalf(
				"\ni:%vgot:%v\nwant:%v\n",
				i, got, e.want,
			)
		}
	}
}

func Benchmark_examples_maxProductSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxProductSort(bench.nums)
	}
}

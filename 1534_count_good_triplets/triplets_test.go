package triplets

import "testing"

var examples = []struct {
	want    int
	arr     []int
	a, b, c int
}{
	{
		want: 4,
		arr:  []int{3, 0, 1, 1, 9, 7},
		a:    7, b: 2, c: 3,
	},
	{
		want: 0,
		arr:  []int{1, 1, 2, 2, 3},
		a:    0, b: 0, c: 1,
	},
}

var bench = examples[0]

func Test_examples_countGoodTriplets(t *testing.T) {
	for i, e := range examples {
		got := countGoodTriplets(e.arr, e.a, e.b, e.c)
		if got != e.want {
			t.Fatalf(
				"\ni:%v\ngot:%v\nwant:%v\narr:%v\n",
				i, got, e.want, e.arr,
			)
		}
	}
}

func Benchmark_examples_countGoodTriplets(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countGoodTriplets(bench.arr, bench.a, bench.b, bench.c)
	}
}

func Test_examples_countGoodTriplets2(t *testing.T) {
	for i, e := range examples {
		got := countGoodTriplets2(e.arr, e.a, e.b, e.c)
		if got != e.want {
			t.Fatalf(
				"\ni:%v\ngot:%v\nwant:%v\narr:%v\n",
				i, got, e.want, e.arr,
			)
		}
	}
}

func Benchmark_examples_countGoodTriplets2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countGoodTriplets2(bench.arr, bench.a, bench.b, bench.c)
	}
}

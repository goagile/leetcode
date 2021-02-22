package sum

import "testing"

var examples = []struct {
	want int
	arr  []int
}{
	{
		want: 58,
		arr:  []int{1, 4, 2, 5, 3},
	},
	{
		want: 3,
		arr:  []int{1, 2},
	},
	{
		want: 66,
		arr:  []int{10, 11, 12},
	},
}

var bench = examples[0]

func Test_examples_sumOddLengthSubarrays(t *testing.T) {
	for i, e := range examples {
		got := sumOddLengthSubarrays(e.arr)
		if got != e.want {
			t.Fatalf(
				"\ni:%v\ngot:%v\nwant:%v\n",
				i, got, e.want,
			)
		}
	}
}

func Benchmark_examples_sumOddLengthSubarrays(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumOddLengthSubarrays(bench.arr)
	}
}

func Test_examples_sumOddLengthSubarraysFun(t *testing.T) {
	for i, e := range examples {
		got := sumOddLengthSubarraysFun(e.arr)
		if got != e.want {
			t.Fatalf(
				"\ni:%v\ngot:%v\nwant:%v\n",
				i, got, e.want,
			)
		}
	}
}

func Benchmark_examples_sumOddLengthSubarraysFun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumOddLengthSubarraysFun(bench.arr)
	}
}

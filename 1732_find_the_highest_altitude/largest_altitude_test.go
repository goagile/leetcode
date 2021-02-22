package largest

import "testing"

var examples = []struct {
	want int
	gain []int
}{
	{
		want: 0,
		gain: []int{-4, -3, -2, -1, 4, 3, 2},
	},
	{
		want: 0,
		gain: []int{},
	},
	{
		want: 1,
		gain: []int{-5, 1, 5, 0, -7},
	},
}

var bench = examples[0]

func Test_examples_largestAltitude(t *testing.T) {
	for i, e := range examples {
		got := largestAltitude(e.gain)
		if got != e.want {
			t.Fatalf(
				"\ni:%v\ngot:%v\nwant:%v\n",
				i, got, e.want,
			)
		}
	}
}

func Benchmark_examples_largestAltitude(b *testing.B) {
	for i := 0; i < b.N; i++ {
		largestAltitude(bench.gain)
	}
}

func Test_examples_largestAltitudeMathMax(t *testing.T) {
	for i, e := range examples {
		got := largestAltitudeMathMax(e.gain)
		if got != e.want {
			t.Fatalf(
				"\ni:%v\ngot:%v\nwant:%v\n",
				i, got, e.want,
			)
		}
	}
}

func Benchmark_examples_largestAltitudeMathMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		largestAltitudeMathMax(bench.gain)
	}
}

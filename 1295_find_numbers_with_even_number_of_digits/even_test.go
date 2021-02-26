package even

import "testing"

var examples = []struct {
	want int
	nums []int
}{
	{
		want: 2,
		nums: []int{12, 345, 2, 6, 7896},
	},
	{
		want: 1,
		nums: []int{555, 901, 482, 1771},
	},
}

var bench = examples[0]

func Test_examples_findNumbers(t *testing.T) {
	for i, e := range examples {
		got := findNumbers(e.nums)
		if got != e.want {
			t.Fatalf(
				"\ni:%v\ngot:%v\nwant:%v\n",
				i, got, e.want,
			)
		}
	}
}

func Benchmark_examples_findNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findNumbers(bench.nums)
	}
}

func Test_examples_findNumbers2(t *testing.T) {
	for i, e := range examples {
		got := findNumbers2(e.nums)
		if got != e.want {
			t.Fatalf(
				"\ni:%v\ngot:%v\nwant:%v\n",
				i, got, e.want,
			)
		}
	}
}

func Benchmark_examples_findNumbers2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findNumbers2(bench.nums)
	}
}

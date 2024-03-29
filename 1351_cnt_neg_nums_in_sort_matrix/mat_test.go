package mat

import "testing"

var examples = []struct {
	want int
	grid [][]int
}{
	{
		want: 8,
		grid: [][]int{
			{4, 3, 2, -1},
			{3, 2, 1, -1},
			{1, 1, -1, -2},
			{-1, -1, -2, -3},
		},
	},
	{
		want: 0,
		grid: [][]int{
			{3, 2},
			{1, 0},
		},
	},
	{
		want: 3,
		grid: [][]int{
			{1, -1},
			{-1, -1},
		},
	},
	{
		want: 1,
		grid: [][]int{
			{-1},
		},
	},
}

var bench = examples[0]

func Test_examples_countNegatives(t *testing.T) {
	for i, e := range examples {
		got := countNegatives(e.grid)
		if got != e.want {
			t.Fatalf(
				"\ni:%v\ngot:%v\nwant:%v\n",
				i, got, e.want,
			)
		}
	}
}

func Benchmark_examples_countNegatives(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countNegatives(bench.grid)
	}
}

func Test_examples_countNegativesBisect(t *testing.T) {
	for i, e := range examples {
		got := countNegativesBisect(e.grid)
		if got != e.want {
			t.Fatalf(
				"\ni:%v\ngot:%v\nwant:%v\n",
				i, got, e.want,
			)
		}
	}
}

func Benchmark_examples_countNegativesBisect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countNegativesBisect(bench.grid)
	}
}

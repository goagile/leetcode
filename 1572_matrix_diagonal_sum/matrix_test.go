package matrix

import "testing"

var examples = []struct {
	want int
	mat  [][]int
}{
	{
		want: 25,
		mat: [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
	},
	{
		want: 8,
		mat: [][]int{
			{1, 1, 1, 1},
			{1, 1, 1, 1},
			{1, 1, 1, 1},
			{1, 1, 1, 1},
		},
	},
	{
		want: 5,
		mat: [][]int{
			{5},
		},
	},
}

var bench = examples[0]

func Test_examples_diagonalSum(t *testing.T) {
	for i, e := range examples {
		got := diagonalSum(e.mat)
		if got != e.want {
			t.Fatalf(
				"\ni:%v\ngot:%v\nwant:%v\n",
				i, got, e.want,
			)
		}
	}
}

func Benchmark_examples_diagonalSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		diagonalSum(bench.mat)
	}
}

func Test_examples_diagonalSumOdd(t *testing.T) {
	for i, e := range examples {
		got := diagonalSumOdd(e.mat)
		if got != e.want {
			t.Fatalf(
				"\ni:%v\ngot:%v\nwant:%v\n",
				i, got, e.want,
			)
		}
	}
}

func Benchmark_examples_diagonalSumOdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		diagonalSumOdd(bench.mat)
	}
}

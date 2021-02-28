package matrix

import "testing"

var examples = []struct {
	want    int
	n, m    int
	indices [][]int
}{
	{
		want: 6,
		n:    2, m: 3,
		indices: [][]int{{0, 1}, {1, 1}},
	},
	{
		want: 0,
		n:    2, m: 2,
		indices: [][]int{{1, 1}, {0, 0}},
	},
}

func Test_examples_oddCells(t *testing.T) {
	for i, e := range examples {
		got := oddCells(e.n, e.m, e.indices)
		if got != e.want {
			t.Fatalf(
				"\ni:%v\ngot:%v\nwant:%v\n",
				i, got, e.want,
			)
		}
	}
}
